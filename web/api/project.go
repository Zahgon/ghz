package api

import (
	"github.com/bojand/ghz/web/model"
	"github.com/labstack/echo"
)

// ProjectDatabase interface for encapsulating database access.
type ProjectDatabase interface {
	CreateProject(project *model.Project) error
	FindProjectByID(id uint) (*model.Project, error)
	UpdateProject(*model.Project) error
	DeleteProject(*model.Project) error
	CountProjects() (uint, error)
	ListProjects(limit, page uint, sortField, order string) ([]*model.Project, error)
}

// The ProjectAPI provides handlers for managing projects.
type ProjectAPI struct {
	DB ProjectDatabase
}

// ProjectList response
type ProjectList struct {
	Total uint             `json:"total"`
	Data  []*model.Project `json:"data"`
}

// CreateProject creates a project
func (api *ProjectAPI) CreateProject(ctx echo.Context) error { _ = "STUB: not implemented"; return nil }

// UpdateProject updates a project
func (api *ProjectAPI) UpdateProject(ctx echo.Context) error { _ = "STUB: not implemented"; return nil }

// GetProject gets a project
func (api *ProjectAPI) GetProject(ctx echo.Context) error { _ = "STUB: not implemented"; return nil }

// ListProjects lists projects
func (api *ProjectAPI) ListProjects(ctx echo.Context) error { _ = "STUB: not implemented"; return nil }

// DeleteProject deletes a project
func (api *ProjectAPI) DeleteProject(ctx echo.Context) error { _ = "STUB: not implemented"; return nil }

func findProject(FindProjectByID func(id uint) (*model.Project, error), ctx echo.Context) (*model.Project, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (api *ProjectAPI) bindAndValidate(ctx echo.Context, p *model.Project) error {
	_ = "STUB: not implemented"
	return nil
}
