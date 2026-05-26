package runner

// RequestCounter gets the request count
type RequestCounter interface {
	Get() uint64
}

// Counter is an implementation of the request counter
type Counter struct {
	c uint64
}

// Get retrieves the current count
func (c *Counter) Get() uint64 { _ = "STUB: not implemented"; return 0 }

// Inc increases the current count
func (c *Counter) Inc() uint64 { _ = "STUB: not implemented"; return 0 }
