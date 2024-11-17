package user

import (
	"net/http"
	"strconv"

	"github.com/arnavsurve/workspaced/pkg/db"
	"github.com/labstack/echo/v4"
)

func HandleGetUserById(c echo.Context, store *db.Store) error {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}
	account, err := store.GetAccountById(intId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, account)
}

