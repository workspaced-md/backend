package user

import (
	"log"
	"net/http"

	"github.com/arnavsurve/workspaced/pkg/db"
	"github.com/arnavsurve/workspaced/pkg/shared"
	"github.com/labstack/echo/v4"
)

func HandleEditUser(c echo.Context, store *db.Store) error {
	account := shared.Account{}
	err := c.Bind(&account)
	if err != nil {
		return err
	}

	// TODO
	// account.Id = c.Get("user").(*jwt.Token).Id

	err = store.EditAccount(&account)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, "Failed to edit account")
	}

	return c.String(http.StatusOK, "Account edited successfully")
}
