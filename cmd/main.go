package main

import (
	"log"

	"github.com/arnavsurve/md/pkg/handlers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env: %v", err)
	}
}

func main() {
	cors := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")
			c.Response().Header().Set(echo.HeaderAccessControlAllowMethods, "GET, POST")
			c.Response().Header().Set(echo.HeaderAccessControlAllowHeaders, "Content-Type")
			return next(c)
		}
	}

	e := echo.New()
 	e.Static("/", "static")
	e.GET("/", func(c echo.Context) error {
		return c.File("./static/src/index.html")
	}, cors)
	e.POST("/upload", handlers.HandleUpload, cors)
	e.GET("/markdown", handlers.HandleMarkdown, cors)
	e.Logger.Fatal(e.Start(":8080"))
}
