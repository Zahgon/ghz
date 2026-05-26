package api

import (
	"github.com/bojand/ghz/web/model"
	"github.com/labstack/echo"
)

// OptionsDatabase interface for encapsulating database access.
type OptionsDatabase interface {
	GetOptionsForReport(uint) (*model.Options, error)
}

// The OptionsAPI provides handlers
type OptionsAPI struct {
	DB OptionsDatabase
}

// GetOptions gets options for a report
func (api *OptionsAPI) GetOptions(ctx echo.Context) error { _ = "STUB: not implemented"; return nil }
