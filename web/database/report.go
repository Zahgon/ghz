package database

import (
	"github.com/bojand/ghz/web/model"
)

// FindReportByID gets the report by id
func (d *Database) FindReportByID(id uint) (*model.Report, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CountReports returns the number of reports
func (d *Database) CountReports() (uint, error) { _ = "STUB: not implemented"; return 0, nil }

// CountReportsForProject returns the number of reports
func (d *Database) CountReportsForProject(pid uint) (uint, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// CreateReport creates a new report
func (d *Database) CreateReport(r *model.Report) error { _ = "STUB: not implemented"; return nil }

// DeleteReport deletes an existing report
func (d *Database) DeleteReport(r *model.Report) error { _ = "STUB: not implemented"; return nil }

// DeleteReportBulk performans a bulk of deletes
func (d *Database) DeleteReportBulk(ids []uint) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// FindPreviousReport find previous report for the report id
func (d *Database) FindPreviousReport(rid uint) (*model.Report, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindLatestReportForProject returns the latest / most recent report for project
func (d *Database) FindLatestReportForProject(pid uint) (*model.Report, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListReports lists reports using sorting
func (d *Database) ListReports(limit, page uint, sortField, order string) ([]*model.Report, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListReportsForProject lists reports using sorting
func (d *Database) ListReportsForProject(pid, limit, page uint, sortField, order string) ([]*model.Report, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Database) listReports(byProject bool, pid, limit, page uint, sortField, order string) ([]*model.Report, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
