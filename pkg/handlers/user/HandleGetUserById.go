package user

import (
	"net/http"
	"strconv"

	"github.com/arnavsurve/workspaced/pkg/auth"
	"github.com/arnavsurve/workspaced/pkg/db"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func HandleGetUserById(c echo.Context, store *db.Store) error {
	// Extract user ID from JWT
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.JwtClaims)
	jwtUserId := claims.UserId

	// Read user ID from URL
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	// Check if user ID from JWT matches user ID from URL
	if intId != jwtUserId {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
	}

	// Fetch user account from database
	account, err := store.GetAccountById(intId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, account)
}
