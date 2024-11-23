package workspace

import (
	"log"
	"net/http"

	"github.com/arnavsurve/workspaced/pkg/auth"
	"github.com/arnavsurve/workspaced/pkg/db"
	"github.com/arnavsurve/workspaced/pkg/shared"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// HandleCreateWorkspace creates a new workspace in the database
func HandleCreateWorkspace(c echo.Context, store *db.Store) error {
	claimsId := c.Get("user").(*jwt.Token).Claims.(*auth.JwtClaims).UserId

	workspace := shared.Workspace{}
	if err := c.Bind(&workspace); err != nil {
		log.Print(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to read body"})
	}

	// Verify user ID in JWT matches the owner ID of the workspace
	if claimsId != workspace.OwnerId {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
	}

	err := store.CreateWorkspace(&workspace)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create workspace"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Workspace created successfully"})
}
