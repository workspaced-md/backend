package workspace

import (
	"log"
	"net/http"
	"strconv"

	"github.com/arnavsurve/workspaced/pkg/auth"
	"github.com/arnavsurve/workspaced/pkg/db"
	"github.com/arnavsurve/workspaced/pkg/shared"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// HandleEditWorkspace updates a workspace's name, description, and is_private fields in the database
func HandleEditWorkspace(c echo.Context, store *db.Store) error {
	workspaceId := c.Param("workspaceId")
	intWorkspaceId, err := strconv.Atoi(workspaceId)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid workspace ID"})
	}

	workspace := shared.Workspace{}
	if err := c.Bind(&workspace); err != nil {
		log.Print(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to read body"})
	}

	// Verify user ID in JWT matches the owner ID of the workspace
	claimsId := c.Get("user").(*jwt.Token).Claims.(*auth.JwtClaims).UserId
	if claimsId != workspace.OwnerId {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
	}

	workspace.Id = intWorkspaceId
	err = store.EditWorkspace(&workspace)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to edit workspace"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Workspace edited successfully"})
}
