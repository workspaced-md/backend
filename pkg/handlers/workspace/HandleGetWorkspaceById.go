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

// HandleGetWorkspaceById retrieves a workspace by its ID
func HandleGetWorkspaceById(c echo.Context, store *db.Store) error {
	workspaceId := c.Param("id")
	intWorkspaceId, err := strconv.Atoi(workspaceId)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid workspace ID"})
	}

	workspace, err := store.GetWorkspaceById(intWorkspaceId)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to get workspace"})
	}

	// Verify user ID in JWT matches the requested user ID
	claimsId := c.Get("user").(*jwt.Token).Claims.(*auth.JwtClaims).UserId
	if claimsId != workspace.OwnerId {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
	}

	return c.JSON(http.StatusOK, workspace)
}
