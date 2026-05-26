package load

import (
	"time"
)

// WorkerTicker is the interface controlling worker parallelism.
type WorkerTicker interface {
	// Ticker returns a channel which sends TickValues
	// When a value is received the number of workers should be appropriately
	// increased or decreased given by the delta property.
	Ticker() <-chan TickValue

	// Run starts the worker ticker
	Run()

	// Finish closes the channel
	Finish()
}

// TickValue is the tick value sent over the ticker channel.
type TickValue struct {
	Delta int  // Delta value representing worker increase or decrease
	Done  bool // A flag representing whether the ticker is done running. Once true no more values should be received over the ticker channel.
}

// ConstWorkerTicker represents a constant number of workers.
// It would send one value for initial number of workers to start.
type ConstWorkerTicker struct {
	C chan TickValue // The tick value channel
	N uint           // The number of workers
}

// Ticker returns the ticker channel.
func (c *ConstWorkerTicker) Ticker() <-chan TickValue {
	_ = "STUB: not implemented"

	// Run runs the ticker.
	return nil
}

func (c *ConstWorkerTicker) Run() { _ = "STUB: not implemented"; return }

// Finish closes the channel.
func (c *ConstWorkerTicker) Finish() {
	_ = "STUB: not implemented"

	// StepWorkerTicker is the worker ticker that implements step adjustments to worker concurrency.
	return
}

type StepWorkerTicker struct {
	C chan TickValue // The tick value channel

	Start        uint          // Starting number of workers
	Step         int           // Step change
	StepDuration time.Duration // Duration to apply the step change
	Stop         uint          // Final number of workers
	MaxDuration  time.Duration // Maximum duration
}

// Ticker returns the ticker channel.
func (c *StepWorkerTicker) Ticker() <-chan TickValue {
	_ = "STUB: not implemented"

	// Run runs the ticker.
	return nil
}

func (c *StepWorkerTicker) Run() { _ = "STUB: not implemented"; return }

// we have load duration and we eclipsed it

// if we have step up and stop value is > current count
// send the final diff

// if we have step down and stop value is < current count
// send the final diff

// send done signal

// we do not have load duration
// if we have stop and are step up and current count >= stop
// or if we have stop and are step down and current count <= stop
// send done signal

// Finish closes the channel.
func (c *StepWorkerTicker) Finish() {
	_ = "STUB: not implemented"

	// LineWorkerTicker is the worker ticker that implements line adjustments to concurrency.
	// Essentially this is same as step worker with 1s step duration.
	return
}

type LineWorkerTicker struct {
	C chan TickValue // The tick value channel

	Start       uint          // Starting number of workers
	Slope       int           // Slope value to adjust the number of workers
	Stop        uint          // Final number of workers
	MaxDuration time.Duration // Maximum adjustment duration

	stepTicker StepWorkerTicker
}

// Ticker returns the ticker channel.
func (c *LineWorkerTicker) Ticker() <-chan TickValue {
	_ = "STUB: not implemented"

	// Run runs the ticker.
	return nil
}

func (c *LineWorkerTicker) Run() { _ = "STUB: not implemented"; return }

// Finish closes the internal tick value channel.
func (c *LineWorkerTicker) Finish() { _ = "STUB: not implemented"; return }
