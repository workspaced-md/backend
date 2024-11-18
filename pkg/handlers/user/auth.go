package user

import (
	"os"
	"time"

	"github.com/arnavsurve/workspaced/pkg/shared"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(account *shared.Account) (string, error) {
	claims := jwt.MapClaims{
		"email": account.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(), // 72 hour expiration
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(secret))
}
