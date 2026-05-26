package model

// Project represents a project
type Project struct {
	Model
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description"`
	Status      Status `json:"status" gorm:"not null"`
}

// BeforeCreate is a GORM hook called when a model is created
func (p *Project) BeforeCreate() error { _ = "STUB: not implemented"; return nil }

// BeforeUpdate is a GORM hook called when a model is updated
func (p *Project) BeforeUpdate() error { _ = "STUB: not implemented"; return nil }

// BeforeSave is a GORM hook called when a model is created or updated
func (p *Project) BeforeSave() error { _ = "STUB: not implemented"; return nil }
