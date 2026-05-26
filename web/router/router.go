package router

import (
	"io"

	"github.com/bojand/ghz/web/api"
	"github.com/bojand/ghz/web/config"
	"github.com/bojand/ghz/web/database"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"

	// for bundled resources
	_ "github.com/bojand/ghz/web/router/statik"
)

// New creates new server
func New(db *database.Database, appInfo *api.ApplicationInfo, conf *config.Config) (*echo.Echo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// API

// Projects

// Reports by Project

// Reports

// Ingest

// Ingest to project

// Info

// Frontend

// load the precompiled statik fs

// get the index file

// wrap the handler

// our custom handler

// if root just pass through to the fs handler

// if it has an extension means it's a file
// so pass through to the fs handler

// otherwise serve the index file
// React router will handle the path from there on

// CustomValidator is our validator for the API
type CustomValidator struct {
	validator *validator.Validate
}

// Validate validates the input
func (cv *CustomValidator) Validate(i interface{}) error { _ = "STUB: not implemented"; return nil }

func getLogLevel(config *config.Config) log.Lvl { _ = "STUB: not implemented"; return *new(log.Lvl) }

func getLogOutput(config *config.Config) (io.Writer, error) {
	_ = "STUB: not implemented"
	return *new(io.Writer), nil
}

// PrintRoutes prints routes in the server
func PrintRoutes(echoServer *echo.Echo) { _ = "STUB: not implemented"; return }
