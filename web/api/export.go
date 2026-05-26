package api

import (
	"time"

	"github.com/alecthomas/template"
	"github.com/bojand/ghz/runner"
	"github.com/bojand/ghz/web/model"
	"github.com/labstack/echo"
)

// ExportDatabase interface for encapsulating database access.
type ExportDatabase interface {
	FindReportByID(uint) (*model.Report, error)
	GetHistogramForReport(uint) (*model.Histogram, error)
	GetOptionsForReport(uint) (*model.Options, error)
	ListAllDetailsForReport(uint) ([]*model.Detail, error)
}

// The ExportAPI provides handlers.
type ExportAPI struct {
	DB ExportDatabase
}

// JSONExportRespose is the response to JSON export
type JSONExportRespose struct {
	model.Report

	Options *model.OptionsInfo `json:"options,omitempty"`

	Histogram model.BucketList `json:"histogram"`

	Details []*runner.ResultDetail `json:"details"`
}

const (
	csvTmpl = `
duration (ms),status,error{{ range $i, $v := . }}
{{ formatDuration .Latency 1000000 }},{{ .Status }},{{ .Error }}{{ end }}
`
)

var tmplFuncMap = template.FuncMap{
	"formatDuration": formatDuration,
}

func formatDuration(duration time.Duration, div int64) string { _ = "STUB: not implemented"; return "" }

// GetExport does export for the report
func (api *ExportAPI) GetExport(ctx echo.Context) error { _ = "STUB: not implemented"; return nil }
