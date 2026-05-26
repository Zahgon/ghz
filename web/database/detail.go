package database

import (
	"github.com/bojand/ghz/web/model"
)

// ListAllDetailsForReport lists all details for report
func (d *Database) ListAllDetailsForReport(rid uint) ([]*model.Detail, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateDetailsBatch creates a batch of details
// Returns the number successfully created, and the number failed
func (d *Database) CreateDetailsBatch(rid uint, s []*model.Detail) (uint, uint) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (d *Database) createDetail(detail *model.Detail) error { _ = "STUB: not implemented"; return nil }
