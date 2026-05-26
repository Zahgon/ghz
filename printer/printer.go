package printer

import (
	"io"
	"time"

	"github.com/alecthomas/template"
	"github.com/bojand/ghz/runner"
)

const (
	barChar = "∎"
)

// ReportPrinter is used for printing the report
type ReportPrinter struct {
	Out    io.Writer
	Report *runner.Report
}

// Print the report using the given format
// If format is "csv" detailed listing is printer in csv format.
// Otherwise the summary of results is printed.
//
// Supported Format:
//
//	summary
//	csv
//	json
//	pretty
//	html
//	influx-summary
//	influx-details
func (rp *ReportPrinter) Print(format string) error { _ = "STUB: not implemented"; return nil }

func (rp *ReportPrinter) print(s string) error { _ = "STUB: not implemented"; return nil }

var tmplFuncMap = template.FuncMap{
	"formatMilli":      formatMilli,
	"formatSeconds":    formatSeconds,
	"histogram":        histogram,
	"jsonify":          jsonify,
	"formatMark":       formatMarkMs,
	"formatPercent":    formatPercent,
	"formatStatusCode": formatStatusCode,
	"formatErrorDist":  formatErrorDist,
	"formatDate":       formatDate,
	"formatNanoUnit":   formatNanoUnit,
}

func jsonify(v interface{}, pretty bool) string { _ = "STUB: not implemented"; return "" }

func formatNanoUnit(d time.Duration) string { _ = "STUB: not implemented"; return "" }

func formatMilli(duration float64) string { _ = "STUB: not implemented"; return "" }

func formatDate(d time.Time) string { _ = "STUB: not implemented"; return "" }

func formatSeconds(duration float64) string { _ = "STUB: not implemented"; return "" }

func formatPercent(num int, total uint64) string { _ = "STUB: not implemented"; return "" }

func histogram(buckets []runner.Bucket) string { _ = "STUB: not implemented"; return "" }

// Normalize bar lengths.

func formatMarkMs(m float64) string { _ = "STUB: not implemented"; return "" }

func formatStatusCode(statusCodeDist map[string]int) string { _ = "STUB: not implemented"; return "" }

// bytes.Buffer can be assumed to not fail on write

// bytes.Buffer can be assumed to not fail on write

func formatErrorDist(errDist map[string]int) string { _ = "STUB: not implemented"; return "" }

// bytes.Buffer can be assumed to not fail on write

// bytes.Buffer can be assumed to not fail on write
