package model

import (
	"database/sql/driver"

	"github.com/bojand/ghz/runner"
	"github.com/jinzhu/gorm"
)

// Options represents a report detail
type Options struct {
	Model

	Report *Report `json:"-"`

	// Run id
	ReportID uint `json:"reportID" gorm:"type:integer REFERENCES reports(id) ON DELETE CASCADE;not null"`

	Info *OptionsInfo `json:"info,omitempty" gorm:"type:TEXT"`
}

// BeforeSave is called by GORM before save
func (o *Options) BeforeSave(scope *gorm.Scope) error { _ = "STUB: not implemented"; return nil }

// OptionsInfo represents the report options
type OptionsInfo runner.Options

// Value converts options struct to a database value
func (o OptionsInfo) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

// Scan converts database value to an Options struct
func (o *OptionsInfo) Scan(src interface{}) error { _ = "STUB: not implemented"; return nil }
