package runner

import (
	"time"
)

// Reporter gathers all the results
type Reporter struct {
	config *RunConfig

	results chan *callResult
	done    chan bool

	totalLatenciesSec float64

	details []ResultDetail

	errorDist      map[string]int
	statusCodeDist map[string]int
	totalCount     uint64
}

// Options represents the request options
// TODO fix casing and consistency
type Options struct {
	Call              string   `json:"call,omitempty"`
	Host              string   `json:"host,omitempty"`
	Proto             string   `json:"proto,omitempty"`
	Protoset          string   `json:"protoset,omitempty"`
	ImportPaths       []string `json:"import-paths,omitempty"`
	EnableCompression bool     `json:"enable-compression,omitempty"`

	CACert    string `json:"cacert,omitempty"`
	Cert      string `json:"cert,omitempty"`
	Key       string `json:"key,omitempty"`
	CName     string `json:"cname,omitempty"`
	SkipTLS   bool   `json:"skipTLS,omitempty"`
	Insecure  bool   `json:"insecure"`
	Authority string `json:"authority,omitempty"`

	RPS              int           `json:"rps,omitempty"`
	LoadSchedule     string        `json:"load-schedule"`
	LoadStart        int           `json:"load-start"`
	LoadEnd          int           `json:"load-end"`
	LoadStep         int           `json:"load-step"`
	LoadStepDuration time.Duration `json:"load-step-duration"`
	LoadMaxDuration  time.Duration `json:"load-max-duration"`

	Concurrency   int           `json:"concurrency,omitempty"`
	CSchedule     string        `json:"concurrency-schedule"`
	CStart        int           `json:"concurrency-start"`
	CEnd          int           `json:"concurrency-end"`
	CStep         int           `json:"concurrency-step"`
	CStepDuration time.Duration `json:"concurrency-step-duration"`
	CMaxDuration  time.Duration `json:"concurrency-max-duration"`

	Total int  `json:"total,omitempty"`
	Async bool `json:"async,omitempty"`

	Connections   int           `json:"connections,omitempty"`
	Duration      time.Duration `json:"duration,omitempty"`
	Timeout       time.Duration `json:"timeout,omitempty"`
	DialTimeout   time.Duration `json:"dial-timeout,omitempty"`
	KeepaliveTime time.Duration `json:"keepalive,omitempty"`

	Data     interface{}        `json:"data,omitempty"`
	Binary   bool               `json:"binary"`
	Metadata *map[string]string `json:"metadata,omitempty"`

	CPUs int    `json:"CPUs"`
	Name string `json:"name,omitempty"`

	SkipFirst   int  `json:"skipFirst,omitempty"`
	CountErrors bool `json:"count-errors,omitempty"`
}

// Report holds the data for the full test
type Report struct {
	Name      string     `json:"name,omitempty"`
	EndReason StopReason `json:"endReason,omitempty"`
	Date      time.Time  `json:"date"`
	Options   Options    `json:"options,omitempty"`

	Count   uint64        `json:"count"`
	Total   time.Duration `json:"total"`
	Average time.Duration `json:"average"`
	Fastest time.Duration `json:"fastest"`
	Slowest time.Duration `json:"slowest"`
	Rps     float64       `json:"rps"`

	ErrorDist      map[string]int `json:"errorDistribution"`
	StatusCodeDist map[string]int `json:"statusCodeDistribution"`

	LatencyDistribution []LatencyDistribution `json:"latencyDistribution"`
	Histogram           []Bucket              `json:"histogram"`
	Details             []ResultDetail        `json:"details"`

	Tags map[string]string `json:"tags,omitempty"`
}

// MarshalJSON is custom marshal for report to properly format the date
func (r Report) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// LatencyDistribution holds latency distribution data
type LatencyDistribution struct {
	Percentage int           `json:"percentage"`
	Latency    time.Duration `json:"latency"`
}

// Bucket holds histogram data
type Bucket struct {
	// The Mark for histogram bucket in seconds
	Mark float64 `json:"mark"`

	// The count in the bucket
	Count int `json:"count"`

	// The frequency of results in the bucket as a decimal percentage
	Frequency float64 `json:"frequency"`
}

// ResultDetail data for each result
type ResultDetail struct {
	Timestamp time.Time     `json:"timestamp"`
	Latency   time.Duration `json:"latency"`
	Error     string        `json:"error"`
	Status    string        `json:"status"`
}

func newReporter(results chan *callResult, c *RunConfig) *Reporter {
	_ = "STUB: not implemented"
	return nil
}

// Run runs the reporter
func (r *Reporter) Run() { _ = "STUB: not implemented"; return }

// Finalize all the gathered data into a final report
func (r *Reporter) Finalize(stopReason StopReason, total time.Duration) *Report {
	_ = "STUB: not implemented"
	return nil
}

func latencies(latencies []float64) []LatencyDistribution { _ = "STUB: not implemented"; return nil }

// since we're dealing with 0th based ranks we need to
// check if ordinal is a whole number that lands on the percentile
// if so adjust accordingly

func histogram(latencies []float64, slowest, fastest float64) []Bucket {
	_ = "STUB: not implemented"
	return nil
}
