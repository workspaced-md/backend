package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func HandleUpload(c echo.Context) error {
	if c.Request().Method != http.MethodPost {
		return c.String(http.StatusBadRequest, "Invalid request method")
	}

	// parse multipart form
	err := c.Request().ParseMultipartForm(10 << 20) // limit file size to 10MB
	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, "Failed to parse form")
	}

	// retrieve file from the form
	file, header, err := c.Request().FormFile("markdownFile")
	if err != nil {
		return c.String(http.StatusBadRequest, "Failed to get file from form")
	}
	defer file.Close()

	// read target directory from the form
	targetDir := c.FormValue("targetDir")
	if targetDir == "" {
		return c.String(http.StatusBadRequest, "Target directory not provided")
	}

	// get root directory from .env
	rootDir := os.Getenv("ROOT_DIR")
	if rootDir == "" {
		return c.String(http.StatusInternalServerError, "Root directory not configured in env")
	}

	// create destination file path
	dstPath := filepath.Join(rootDir, targetDir, header.Filename)
	dstFile, err := os.Create(dstPath)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, "Failed to create file, possibly an invalid path")
	}
	defer dstFile.Close()

	// copy uploaded file to the destination
	_, err = io.Copy(dstFile, file)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to save file")
	}

	return c.String(http.StatusOK, "File uploaded successfully")
}
