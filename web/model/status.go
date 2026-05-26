package model

// Status represents a status of a project or record
type Status string

// StatusFromString creates a Status from a string
func StatusFromString(str string) Status { _ = "STUB: not implemented"; return *new(Status) }

const (
	// StatusOK means the latest run in test was within the threshold
	StatusOK = Status("ok")

	// StatusFail means the latest run in test was not within the threshold
	StatusFail = Status("fail")
)
