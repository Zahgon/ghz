package runner

import (
	"context"
	"sync"
	"time"

	"github.com/bojand/ghz/load"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic/grpcdynamic"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"

	// To register the xds resolvers and balancers.
	_ "google.golang.org/grpc/xds"
)

// Max size of the buffer of result channel.
const maxResult = 1000000

// result of a call
type callResult struct {
	err       error
	status    string
	duration  time.Duration
	timestamp time.Time
}

// Requester is used for doing the requests
type Requester struct {
	conns    []*grpc.ClientConn
	stubs    []grpcdynamic.Stub
	handlers []*statsHandler

	mtd      *desc.MethodDescriptor
	reporter *Reporter

	config *RunConfig

	results chan *callResult
	stopCh  chan bool
	start   time.Time

	dataProvider     DataProviderFunc
	metadataProvider MetadataProviderFunc

	lock       sync.Mutex
	stopReason StopReason
	workers    []*Worker
}

// NewRequester creates a new requestor from the passed RunConfig
func NewRequester(c *RunConfig) (*Requester, error) { _ = "STUB: not implemented"; return nil, nil }

// use reflection to get method descriptor

// temporary connection for reflection, do not store as requester connections

// purposefully ignoring error as we do not care if there
// is an error on close

// cancel is ignored here as connection.Close() is used.
// See https://godoc.org/google.golang.org/grpc#DialContext

// fill in the rest

// Run makes all the requests and returns a report of results
// It blocks until all work is done.
func (b *Requester) Run() (*Report, error) { _ = "STUB: not implemented"; return nil, nil }

// create a client stub for each connection

// Stop stops the test
func (b *Requester) Stop(reason StopReason) { _ = "STUB: not implemented"; return }

// Finish finishes the test run
func (b *Requester) Finish() *Report { _ = "STUB: not implemented"; return nil }

// Wait until the reporter is done.

func (b *Requester) openClientConns() ([]*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Requester) closeClientConns() { _ = "STUB: not implemented"; return }

func (b *Requester) newClientConn(withStatsHandler bool) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// increase max receive and send message sizes

// cancel is ignored here as connection.Close() is used.
// See https://godoc.org/google.golang.org/grpc#DialContext

// create client connection

func (b *Requester) runWorkers(wt load.WorkerTicker, p load.Pacer) error {
	_ = "STUB: not implemented"
	return nil
}

// worker control ticker goroutine

// increment worker id

// increment connection counter

// wrap around connections if needed

func min(a, b int) int { _ = "STUB: not implemented"; return 0 }

func createWorkerTicker(config *RunConfig) load.WorkerTicker {
	_ = "STUB: not implemented"
	return *new(load.WorkerTicker)
}

func createPacer(config *RunConfig) load.Pacer { _ = "STUB: not implemented"; return *new(load.Pacer) }

func checkState(conn *grpc.ClientConn, states ...connectivity.State) bool {
	_ = "STUB: not implemented"
	return false
}

func connectionOnState(ctx context.Context, conn *grpc.ClientConn, states ...connectivity.State) <-chan bool {
	_ = "STUB: not implemented"
	return nil
}
