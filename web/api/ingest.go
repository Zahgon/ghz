package api

import (
	"github.com/bojand/ghz/runner"
	"github.com/bojand/ghz/web/model"
	"github.com/labstack/echo"
)

// IngestDatabase interface for encapsulating database access.
type IngestDatabase interface {
	CreateProject(*model.Project) error
	CreateReport(*model.Report) error
	CreateHistogram(*model.Histogram) error
	CreateOptions(*model.Options) error
	FindProjectByID(uint) (*model.Project, error)
	FindLatestReportForProject(uint) (*model.Report, error)
	CreateDetailsBatch(uint, []*model.Detail) (uint, uint)
	UpdateProjectStatus(uint, model.Status) error
}

// IngestResponse is the response to the ingest endpoint
type IngestResponse struct {
	// Created project
	Project *model.Project `json:"project"`

	// Created report
	Report *model.Report `json:"report"`

	// Created Options
	Options *model.Options `json:"options"`

	// Created Histogram
	Histogram *model.Histogram `json:"histogram"`

	// The summary of created details
	Details *DetailsCreated `json:"details"`
}

// DetailsCreated summary of how many details got created and how many failed
type DetailsCreated struct {
	// Number of successfully created detail objects
	Success uint `json:"success"`

	// Number of failed detail objects
	Fail uint `json:"fail"`
}

// The IngestAPI provides handlers for ingesting and processing reports.
type IngestAPI struct {
	DB IngestDatabase
}

// IngestRequest is the raw report
type IngestRequest runner.Report

// Ingest creates data from raw report
func (api *IngestAPI) Ingest(ctx echo.Context) error { _ = "STUB: not implemented"; return nil }

// Project

// IngestToProject ingests data into a specific project
func (api *IngestAPI) IngestToProject(ctx echo.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (api *IngestAPI) ingestToProject(p *model.Project, ir *IngestRequest, ctx echo.Context) error {
	_ = "STUB: not implemented"

	// first get latest (we'll need it later)
	return nil
}

// Report

// Options

// Histogram

// Details

// Update project status if needed

// Response

func convertIngestToReport(pid uint, ir *IngestRequest) *model.Report {
	_ = "STUB: not implemented"
	return nil
}

// status

func bindAndValidateInput(ctx echo.Context, ir *IngestRequest) error {
	_ = "STUB: not implemented"
	return nil
}
