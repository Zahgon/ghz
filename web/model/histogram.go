package model

import (
	"database/sql/driver"

	"github.com/bojand/ghz/runner"
	"github.com/jinzhu/gorm"
)

// BucketList is a slice of buckets
type BucketList []*runner.Bucket

// Value converts struct to a database value
func (bl BucketList) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

// Scan converts database value to a struct
func (bl *BucketList) Scan(src interface{}) error { _ = "STUB: not implemented"; return nil }

// Histogram represents a histogram
type Histogram struct {
	Model

	ReportID uint    `json:"reportID" gorm:"type:integer REFERENCES reports(id) ON DELETE CASCADE;not null"`
	Report   *Report `json:"-"`

	Buckets BucketList `json:"buckets" gorm:"type:TEXT"`
}

// BeforeSave is called by GORM before save
func (h *Histogram) BeforeSave(scope *gorm.Scope) error { _ = "STUB: not implemented"; return nil }
