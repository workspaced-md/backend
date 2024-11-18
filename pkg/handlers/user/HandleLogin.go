package user

import (
	"log"
	"net/http"

	"github.com/arnavsurve/workspaced/pkg/db"
	"github.com/arnavsurve/workspaced/pkg/shared"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func HandleLogin(c echo.Context, store *db.Store) error {
	req := shared.Account{}
	if err := c.Bind(&req); err != nil {
		log.Printf("Failed to read body: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to read body"})
	}

	// Look up requested user
	account, err := store.GetAccountByEmail(req.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "User not found"})
	}

	// Check password
	if err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(req.Password)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid password"})
	}

	// Generate JWT
	token, err := GenerateJWT(account)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to log in"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User logged in successfully", "token": token})
}
