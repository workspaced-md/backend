package user

import (
	"log"
	"net/http"

	"github.com/arnavsurve/md/pkg/db"
	"github.com/arnavsurve/md/pkg/shared"
	"github.com/labstack/echo/v4"
)

func HandleNewUser(c echo.Context, store *db.Store) error {
	account := shared.Account{}
	err := c.Bind(&account)
	if err != nil {
		return err
	}

	err = store.CreateAccount(&account)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, "Failed to create account")
	}
	
	return c.String(http.StatusOK, "Account created successfully")
}