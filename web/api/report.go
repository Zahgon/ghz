package api

import (
	"github.com/bojand/ghz/web/model"
	"github.com/labstack/echo"
)

// ReportDatabase interface for encapsulating database access.
type ReportDatabase interface {
	CountReports() (uint, error)
	CountReportsForProject(uint) (uint, error)
	FindReportByID(uint) (*model.Report, error)
	FindPreviousReport(uint) (*model.Report, error)
	DeleteReport(*model.Report) error
	DeleteReportBulk([]uint) (int, error)
	ListReports(limit, page uint, sortField, order string) ([]*model.Report, error)
	ListReportsForProject(pid, limit, page uint, sortField, order string) ([]*model.Report, error)
}

// The ReportAPI provides handlers for managing reports.
type ReportAPI struct {
	DB ReportDatabase
}

// ReportList response
type ReportList struct {
	Total uint            `json:"total"`
	Data  []*model.Report `json:"data"`
}

// DeleteReportBulkRequest is the request to delete bulk reports
type DeleteReportBulkRequest struct {
	IDs []uint `json:"ids"`
}

// ListReportsForProject lists reports for a project
func (api *ReportAPI) ListReportsForProject(ctx echo.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// ListReportsAll gets a list of all reports
func (api *ReportAPI) ListReportsAll(ctx echo.Context) error { _ = "STUB: not implemented"; return nil }

func (api *ReportAPI) listReports(forProject bool, projectID uint, ctx echo.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// GetReport gets a report
func (api *ReportAPI) GetReport(ctx echo.Context) error { _ = "STUB: not implemented"; return nil }

// DeleteReport deletes a report
func (api *ReportAPI) DeleteReport(ctx echo.Context) error { _ = "STUB: not implemented"; return nil }

// DeleteReportBulk deletes bulk reports
func (api *ReportAPI) DeleteReportBulk(ctx echo.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// GetPreviousReport gets a previous report
func (api *ReportAPI) GetPreviousReport(ctx echo.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func getReportID(ctx echo.Context) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }
