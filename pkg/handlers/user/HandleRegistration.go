package user

import (
	"log"
	"net/http"

	"github.com/arnavsurve/workspaced/pkg/db"
	"github.com/arnavsurve/workspaced/pkg/shared"
	"github.com/labstack/echo/v4"
)

func HandleNewUser(c echo.Context, store *db.Store) error {
	account := shared.Account{}
	err := c.Bind(&account)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to create user: " + err.Error()})
	}

	err = store.CreateAccount(&account)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to create user: " + err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User created successfully"})
}
