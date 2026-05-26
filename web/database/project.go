package database

import (
	"github.com/bojand/ghz/web/model"
)

// FindProjectByID gets the project by id
func (d *Database) FindProjectByID(id uint) (*model.Project, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CountProjects returns the number of projects
func (d *Database) CountProjects() (uint, error) { _ = "STUB: not implemented"; return 0, nil }

// CreateProject creates a new project
func (d *Database) CreateProject(p *model.Project) error { _ = "STUB: not implemented"; return nil }

// UpdateProject update a project
func (d *Database) UpdateProject(p *model.Project) error { _ = "STUB: not implemented"; return nil }

// DeleteProject deletas an existing project
func (d *Database) DeleteProject(p *model.Project) error { _ = "STUB: not implemented"; return nil }

// UpdateProjectStatus updates the project's status
func (d *Database) UpdateProjectStatus(pid uint, status model.Status) error {
	_ = "STUB: not implemented"
	return nil
}

// use UpdateColumn to circumvent update hooks and not modify updated at time

// ListProjects lists projects using sorting
func (d *Database) ListProjects(limit, page uint, sortField, order string) ([]*model.Project, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
