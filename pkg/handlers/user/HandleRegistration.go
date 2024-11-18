package user

import (
	"log"
	"net/http"

	"github.com/arnavsurve/workspaced/pkg/db"
	"github.com/arnavsurve/workspaced/pkg/shared"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type requestData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HandleNewUser(c echo.Context, store *db.Store) error {
	if c.Bind(&requestData{}) != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to read body"})
	}

	// look up requested user
	account := shared.Account{}
	store.DB.Query("SELECT * FROM accounts WHERE email = $1", requestData.Email).Scan(&account)

	err := c.Bind(&account)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to create user: " + err.Error()})
	}

	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to create user: " + err.Error()})
	}
	account.Password = string(hash)

	err = store.CreateAccount(&account)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to create user: " + err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User created successfully"})
}
