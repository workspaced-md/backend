package user

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/arnavsurve/workspaced/pkg/db"
	"github.com/arnavsurve/workspaced/pkg/shared"
	"github.com/golang-jwt/jwt/v5"
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
	token, err := generateJWT(account)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to create user"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User created successfully", "token": token})
}

func generateJWT(account *shared.Account) (string, error) {
	claims := jwt.MapClaims{
		"email": account.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(), // 72 hour expiration
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(secret))
}
