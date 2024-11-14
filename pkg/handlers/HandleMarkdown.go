package handlers

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
)

// markdownHandler reads a markdown file and renders it as HTML.
func HandleMarkdown(w http.ResponseWriter, r *http.Request) {
	rootDir := os.Getenv("ROOT_DIR")
	if rootDir == "" {
		http.Error(w, "Root directory not set", http.StatusInternalServerError)
		return
	}

	file := r.URL.Query().Get("file")
	if file == "" {
		http.Error(w, "File parameter is missing", http.StatusBadRequest)
		return
	}

	// clean and construct the full file path
	filePath := filepath.Join(rootDir, filepath.Clean(file))
	if !strings.HasPrefix(filepath.Clean(filePath), filepath.Clean(rootDir)) {
		http.Error(w, "Invalid file path", http.StatusBadRequest)
		return
	}

	// read markdown file
	mdContent, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// convert markdown content to HTML
	var buf bytes.Buffer
	markdown := goldmark.New()
	if err := markdown.Convert(mdContent, &buf); err != nil {
		http.Error(w, "Failed to convert markdown", http.StatusInternalServerError)
		return
	}

	// set response headers and write HTML content
	w.Header().Set("Content-Type", "text/html")
	_, err = w.Write(buf.Bytes())
	if err != nil {
		http.Error(w, "Error writing HTML content", http.StatusInternalServerError)
		return
	}
}
