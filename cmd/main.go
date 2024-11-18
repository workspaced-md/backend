package main

import (
	"log"

	"github.com/arnavsurve/workspaced/pkg/db"
	"github.com/arnavsurve/workspaced/pkg/handlers"
	"github.com/arnavsurve/workspaced/pkg/handlers/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	store, err := db.NewStore()
	if err != nil {
		log.Fatal(err)
	}

	store.InitAccountsTable()

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderContentType, echo.HeaderAuthorization},
	}))

	e.Static("/", "static")
	e.POST("/upload", handlers.HandleUpload)
	e.GET("/markdown", handlers.HandleMarkdown)

	userGroup := e.Group("/user")
	userGroup.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderContentType, echo.HeaderAuthorization},
	}))

	userGroup.POST("/register", func(c echo.Context) error {
		return user.HandleNewUser(c, store)
	})
	userGroup.GET("/:id", func(c echo.Context) error {
		return user.HandleGetUserById(c, store)
	})
	userGroup.GET("/:email", func(c echo.Context) error {
		return user.HandleGetUserByEmail(c, store)
	})
	userGroup.PUT("/:id", func(c echo.Context) error {
		return user.HandleEditUser(c, store)
	})
	// TODO /user/login

	e.Logger.Fatal(e.Start(":8080"))
}
