package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arnavsurve/md/pkg/handlers"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env: %v", err)
	}
}

func main() {
	http.HandleFunc("/", handlers.HandleMarkdown)
	http.HandleFunc("/upload", handlers.HandleUpload)
	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
