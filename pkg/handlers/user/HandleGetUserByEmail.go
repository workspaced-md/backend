package user

import (
	"net/http"

	"github.com/arnavsurve/md/pkg/db"
	"github.com/labstack/echo/v4"
)

func HandleGetUserByEmail(c echo.Context, store *db.Store) error {
	email := c.Param("email")
	if email == "" {
		return c.JSON(http.StatusBadRequest, "Invalid email")
	}

	account, err := store.GetAccountByEmail(email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,  err.Error())
	}
	
	return c.JSON(http.StatusOK, account)
}