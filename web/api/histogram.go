package api

import (
	"github.com/bojand/ghz/web/model"
	"github.com/labstack/echo"
)

// HistogramDatabase interface for encapsulating database access.
type HistogramDatabase interface {
	GetHistogramForReport(uint) (*model.Histogram, error)
}

// The HistogramAPI provides handlers.
type HistogramAPI struct {
	DB HistogramDatabase
}

// GetHistogram gets a histogram for the report
func (api *HistogramAPI) GetHistogram(ctx echo.Context) error {
	_ = "STUB: not implemented"
	return nil
}
