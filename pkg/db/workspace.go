package db

import (
	"fmt"

	"github.com/arnavsurve/workspaced/pkg/shared"
)

// CreateWorkspace creates a new workspace in the database.
func (s *Store) CreateWorkspace(workspace *shared.Workspace) error {
	result := s.DB.Create(workspace)
	if result.Error != nil {
		return fmt.Errorf("error creating workspace: %v", result.Error)
	}
	return nil
}

// GetWorkspaceById returns the workspace with the given ID.
func (s *Store) GetWorkspaceById(id int) (*shared.Workspace, error) {
	workspace := &shared.Workspace{}
	result := s.DB.First(workspace, id)
	if result.Error != nil {
		return nil, fmt.Errorf("error getting workspace: %v", result.Error)
	}
	return workspace, nil
}

// GetWorkspacesByAccountId returns all workspaces that belong to the account with the given ID.
func (s *Store) GetWorkspacesByAccountId(accountId int) ([]shared.Workspace, error) {
	workspaces := []shared.Workspace{}
	result := s.DB.Where("owner_id = ?", accountId).Find(&workspaces)
	if result.Error != nil {
		return nil, fmt.Errorf("error getting workspaces: %v", result.Error)
	}
	return workspaces, nil
}

// EditWorkspace updates the workspace with the ID of the given workspace. It updates the name, description, and is_public fields.
func (s *Store) EditWorkspace(workspace *shared.Workspace) error {
	result := s.DB.Model(&shared.Workspace{}).Where("id = ?", workspace.Id).Select("name", "description", "is_public").Updates(workspace)
	if result.Error != nil {
		return fmt.Errorf("error updating workspace: %v", result.Error)
	}
	return nil
}
