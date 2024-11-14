package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/yuin/goldmark"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env: %v", err)
	}
}

// markdownHandler reads a markdown file and renders it as HTML.
func markdownHandler(w http.ResponseWriter, r *http.Request) {
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

	// Clean the file path and construct the full file path
	filePath := filepath.Join(rootDir, filepath.Clean(file))
	if !strings.HasPrefix(filepath.Clean(filePath), filepath.Clean(rootDir)) {
		http.Error(w, "Invalid file path", http.StatusBadRequest)
		return
	}

	// Read the markdown file
	mdContent, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Convert the markdown content to HTML
	var buf bytes.Buffer
	markdown := goldmark.New()
	if err := markdown.Convert(mdContent, &buf); err != nil {
		http.Error(w, "Failed to convert markdown", http.StatusInternalServerError)
		return
	}

	// Set the response headers and write the HTML content
	w.Header().Set("Content-Type", "text/html")
	w.Write(buf.Bytes())
}

func main() {
	http.HandleFunc("/", markdownHandler)
	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
