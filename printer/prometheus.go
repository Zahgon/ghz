package printer

import (
	promtypes "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
)

// https://github.com/prometheus/docs/blob/master/content/docs/instrumenting/exposition_formats.md

func (rp *ReportPrinter) printPrometheus() error { _ = "STUB: not implemented"; return nil }

// histogram

// latency distribution

// errors

func (rp *ReportPrinter) printPrometheusMetricGauge(
	encoder expfmt.Encoder, labels []*promtypes.LabelPair,
	name string, value *promtypes.Gauge) error {
	_ = "STUB: not implemented"
	return nil
}

func (rp *ReportPrinter) getCommonPrometheusLabels() ([]*promtypes.LabelPair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ptrUint64(v uint64) *uint64 { _ = "STUB: not implemented"; return nil }

func ptrFloat64(v float64) *float64 { _ = "STUB: not implemented"; return nil }

func ptrString(v string) *string { _ = "STUB: not implemented"; return nil }

func ptrBoolToStr(v bool) *string { _ = "STUB: not implemented"; return nil }
