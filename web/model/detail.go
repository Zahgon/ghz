package model

import (
	"github.com/bojand/ghz/runner"
)

// Detail represents a report detail
type Detail struct {
	Model

	Report *Report `json:"-"`

	// Run id
	ReportID uint `json:"reportID" gorm:"type:integer REFERENCES reports(id) ON DELETE CASCADE;not null"`

	runner.ResultDetail
}

const layoutISO string = "2006-01-02T15:04:05.000Z"
const layoutISO2 string = "2006-01-02T15:04:05-0700"

// UnmarshalJSON for Detail
func (d *Detail) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// BeforeSave is called by GORM before save
func (d *Detail) BeforeSave() error { _ = "STUB: not implemented"; return nil }
