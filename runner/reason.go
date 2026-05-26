package runner

// StopReason is a reason why the run ended
type StopReason string

// String() is the string representation of threshold
func (s StopReason) String() string { _ = "STUB: not implemented"; return "" }

// UnmarshalJSON prases a Threshold value from JSON string
func (s *StopReason) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON formats a Threshold value into a JSON string
func (s StopReason) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// ReasonFromString creates a Status from a string
func ReasonFromString(str string) StopReason { _ = "STUB: not implemented"; return *new(StopReason) }

const (
	// ReasonNormalEnd indicates a normal end to the run
	ReasonNormalEnd = StopReason("normal")

	// ReasonCancel indicates end due to cancellation
	ReasonCancel = StopReason("cancel")

	// ReasonTimeout indicates run ended due to Z parameter timeout
	ReasonTimeout = StopReason("timeout")
)
