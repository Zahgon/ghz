package load

import (
	"sync"
	"time"
)

// nano is the const for number of nanoseconds in a second
const nano = 1e9

// Pacer defines the interface to control the rate of hit.
type Pacer interface {
	// Pace returns the duration the attacker should wait until
	// making next hit, given an already elapsed duration and
	// completed hits. If the second return value is true, an attacker
	// should stop sending hits.
	Pace(elapsed time.Duration, hits uint64) (wait time.Duration, stop bool)

	// Rate returns a Pacer's instantaneous hit rate (per seconds)
	// at the given elapsed duration of an attack.
	Rate(elapsed time.Duration) float64
}

// A ConstantPacer defines a constant rate of hits.
type ConstantPacer struct {
	Freq uint64 // Frequency of hits per second
	Max  uint64 // Optional maximum allowed hits
}

// String returns a pretty-printed description of the ConstantPacer's behaviour:
//
//	ConstantPacer{Freq: 1} => Constant{1 hits / 1s}
func (cp *ConstantPacer) String() string { _ = "STUB: not implemented"; return "" }

// Pace determines the length of time to sleep until the next hit is sent.
func (cp *ConstantPacer) Pace(elapsed time.Duration, hits uint64) (time.Duration, bool) {
	_ = "STUB: not implemented"
	return *new(time.Duration), false
}

// Zero value = infinite rate

// Running behind, send next hit immediately.

// We would overflow delta if we continued, so stop the attack.

// Zero or negative durations cause time.Sleep to return immediately.

// Rate returns a ConstantPacer's instantaneous hit rate (i.e. requests per second)
// at the given elapsed duration of an attack. Since it's constant, the return
// value is independent of the given elapsed duration.
func (cp *ConstantPacer) Rate(elapsed time.Duration) float64 { _ = "STUB: not implemented"; return 0 }

// hitsPerNs returns the rate in fractional hits per nanosecond.
func (cp *ConstantPacer) hitsPerNs() float64 { _ = "STUB: not implemented"; return 0 }

// StepPacer paces an attack by starting at a given request rate
// and increasing or decreasing with steps at a given step interval and duration.
type StepPacer struct {
	Start        ConstantPacer // Constant start rate
	Step         int64         // Step value
	StepDuration time.Duration // Step duration
	Stop         ConstantPacer // Optional constant stop value
	LoadDuration time.Duration // Optional maximum load duration
	Max          uint64        // Optional maximum allowed hits

	once     sync.Once
	init     bool // TOOO improve this
	constAt  time.Duration
	baseHits uint64
}

func (p *StepPacer) initialize() { _ = "STUB: not implemented"; return }

// Pace determines the length of time to sleep until the next hit is sent.
func (p *StepPacer) Pace(elapsed time.Duration, hits uint64) (time.Duration, bool) {
	_ = "STUB: not implemented"
	return *new(time.Duration), false
}

// Running behind, send next hit immediately.

// const part

// We would overflow wait if we continued, so stop the attack.

// if wait > nano {
// 	intervals := elapsed / nano
// 	wait = (intervals+1)*nano - elapsed
// }

// Rate returns a StepPacer's instantaneous hit rate (i.e. requests per second)
// at the given elapsed duration.
func (p *StepPacer) Rate(elapsed time.Duration) float64 { _ = "STUB: not implemented"; return 0 }

// hits returns the number of hits that have been sent at elapsed duration t.
func (p *StepPacer) hits(t time.Duration) float64 { _ = "STUB: not implemented"; return 0 }

// first step

// previous steps: 1...n

// current step

// String returns a pretty-printed description of the StepPacer's behaviour:
//
//	StepPacer{Step: 1, StepDuration: 5s} => Step{Step:1 hits / 5s}
func (p *StepPacer) String() string { _ = "STUB: not implemented"; return "" }

// LinearPacer paces the hit rate by starting at a given request rate
// and increasing linearly with the given slope at 1s interval.
type LinearPacer struct {
	Start        ConstantPacer // Constant start rate
	Slope        int64         // Slope value to change the rate
	Stop         ConstantPacer // Constant stop rate
	LoadDuration time.Duration // Total maximum load duration
	Max          uint64        // Maximum number of hits

	once sync.Once
	sp   StepPacer
}

// initializes the wrapped step pacer
func (p *LinearPacer) initialize() { _ = "STUB: not implemented"; return }

// Pace determines the length of time to sleep until the next hit is sent.
func (p *LinearPacer) Pace(elapsed time.Duration, hits uint64) (time.Duration, bool) {
	_ = "STUB: not implemented"
	return *new(time.Duration), false
}

// Rate returns a LinearPacer's instantaneous hit rate (i.e. requests per second)
// at the given elapsed duration.
func (p *LinearPacer) Rate(elapsed time.Duration) float64 { _ = "STUB: not implemented"; return 0 }

// String returns a pretty-printed description of the LinearPacer's behaviour:
//
//	LinearPacer{Slope: 1} => Linear{1 hits / 1s}
func (p *LinearPacer) String() string { _ = "STUB: not implemented"; return "" }
