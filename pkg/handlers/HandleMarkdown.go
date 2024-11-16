package handlers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

type markdown struct {
	Content string `json:"content"`
}

func HandleMarkdown(c echo.Context) error {
	rootDir := os.Getenv("ROOT_DIR")
	if rootDir == "" {
		return c.JSON(http.StatusInternalServerError, "Root directory not set")
	}

	file := c.QueryParam("file")
	if file == "" {
		return c.JSON(http.StatusBadRequest, "File parameter is missing")
	}

	// clean and construct the full file path
	filePath := filepath.Join(rootDir, filepath.Clean(file))
	if !strings.HasPrefix(filepath.Clean(filePath), filepath.Clean(rootDir)) {
		return c.JSON(http.StatusBadRequest, "Invalid file path")
	}

	// ensure only .md files are served
	if filepath.Ext(filePath) != ".md" {
		return c.JSON(http.StatusBadRequest, "Invalid file type")
	}

	// read markdown file
	mdContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Failed to read file - does it exist?")
	}

	md := markdown{
		Content: string(mdContent),
	}

	return c.JSON(http.StatusOK, md)
}
