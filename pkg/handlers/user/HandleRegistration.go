package user

import (
	"log"
	"net/http"

	"github.com/arnavsurve/workspaced/pkg/db"
	"github.com/arnavsurve/workspaced/pkg/shared"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func HandleNewUser(c echo.Context, store *db.Store) error {
	req := shared.Account{}
	if err := c.Bind(&req); err != nil {
		log.Printf("Failed to read body: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to read body"})
	}

	// Look up requested user
	account, err := store.GetAccountByEmail(req.Email)
	if err == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "User already exists"})
	}

	// Create a new account
	account = &shared.Account{
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to create user"})
	}
	account.Password = string(hash)

	// Save account to database
	if err = store.CreateAccount(account); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to create user"})
	}

	// Generate JWT
	token, err := GenerateJWT(account)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to create user"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User created successfully", "token": token})
}
