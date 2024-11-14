package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		return
	}

	// parse multipart form
	err := r.ParseMultipartForm(10 << 20) // limit file size to 10MB
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		log.Println(err)
		return
	}

	// retrieve file from the form
	file, header, err := r.FormFile("markdownFile")
	if err != nil {
		http.Error(w, "Failed to get file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// get root directory from .env
	rootDir := os.Getenv("ROOT_DIR")
	if rootDir == "" {
		http.Error(w, "Root directory not configured in env", http.StatusInternalServerError)
		return
	}

	// create destination file path
	dstPath := filepath.Join(rootDir, header.Filename)
	dstFile, err := os.Create(dstPath)
	if err != nil {
		http.Error(w, "Failer to create file", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer dstFile.Close()

	// copy uploaded file to the destination
	_, err = io.Copy(dstFile, file)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s uploaded successfully!", header.Filename)
}
