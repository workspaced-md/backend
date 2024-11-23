package workspace

import (
	"log"
	"net/http"
	"strconv"

	"github.com/arnavsurve/workspaced/pkg/auth"
	"github.com/arnavsurve/workspaced/pkg/db"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// HandleGetWorkspacesByAccountId retrieves all workspaces owned by a user
func HandleGetWorkspacesByAccountId(c echo.Context, store *db.Store) error {
	userId := c.Param("id")
	intUserId, err := strconv.Atoi(userId)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid user ID"})
	}

	// Verify user ID in JWT matches the requested user ID
	claimsId := c.Get("user").(*jwt.Token).Claims.(*auth.JwtClaims).UserId
	if claimsId != intUserId {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
	}

	workspaces, err := store.GetWorkspacesByAccountId(intUserId)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to get workspaces"})
	}

	return c.JSON(http.StatusOK, workspaces)
}
