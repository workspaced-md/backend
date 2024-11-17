package main

import (
	"log"

	"github.com/arnavsurve/md/pkg/db"
	"github.com/arnavsurve/md/pkg/handlers"
	"github.com/arnavsurve/md/pkg/handlers/user"
	"github.com/labstack/echo/v4"
)

func main() {
	store, err := db.NewStore()
	if err != nil {
		log.Fatal(err)
	}

	store.InitAccountsTable()

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
	e.POST("/upload", handlers.HandleUpload, cors)
	e.GET("/markdown", handlers.HandleMarkdown, cors)

	userGroup := e.Group("/user", cors)
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



	e.Logger.Fatal(e.Start(":8080"))
}
