package user

import (
	"net/http"

	"github.com/arnavsurve/workspaced/pkg/db"
	"github.com/labstack/echo/v4"
)

func HandleGetUserByEmail(c echo.Context, store *db.Store) error {
	email := c.Param("email")
	if email == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Email is required"})
	}

	account, err := store.GetAccountByEmail(email)
	account.Password = ""
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to get user: " + err.Error()})
	}

	return c.JSON(http.StatusOK, account)
}
