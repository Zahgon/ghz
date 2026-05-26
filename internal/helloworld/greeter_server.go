package helloworld

import (
	"sync"

	context "golang.org/x/net/context"
	"google.golang.org/grpc/stats"
)

// CallType represents one of the gRPC call types:
// unary, client streaming, server streaming, bidi
type CallType string

// Unary is a unary call
var Unary CallType = "unary"

// ClientStream is a client streaming call
var ClientStream CallType = "cs"

// ServerStream is a server streaming call
var ServerStream CallType = "ss"

// Bidi is a bidi / duplex call
var Bidi CallType = "bidi"

// Greeter implements the GreeterServer for tests
type Greeter struct {
	StreamData []*HelloReply

	Stats *HWStatsHandler

	mutex      *sync.RWMutex
	callCounts map[CallType]int
	calls      map[CallType][][]*HelloRequest

	sendMutex  *sync.RWMutex
	sendCounts map[CallType]map[int]int
}

func randomSleep(max int) { _ = "STUB: not implemented"; return }

func (s *Greeter) recordCall(ct CallType) int { _ = "STUB: not implemented"; return 0 }

func (s *Greeter) recordMessage(ct CallType, callIdx int, msg *HelloRequest) {
	_ = "STUB: not implemented"
	return
}

func (s *Greeter) recordStreamSendCounter(ct CallType, callIdx int) {
	_ = "STUB: not implemented"
	return
}

// SayHello implements helloworld.GreeterServer
func (s *Greeter) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SayHellos lists all hellos
func (s *Greeter) SayHellos(req *HelloRequest, stream Greeter_SayHellosServer) error {
	_ = "STUB: not implemented"
	return nil
}

// SayHelloCS is client streaming handler
func (s *Greeter) SayHelloCS(stream Greeter_SayHelloCSServer) error {
	_ = "STUB: not implemented"
	return nil
}

// SayHelloBidi duplex call handler
func (s *Greeter) SayHelloBidi(stream Greeter_SayHelloBidiServer) error {
	_ = "STUB: not implemented"
	return nil
}

// ResetCounters resets the call counts
func (s *Greeter) ResetCounters() { _ = "STUB: not implemented"; return }

// GetCount gets the count for specific call type
func (s *Greeter) GetCount(key CallType) int { _ = "STUB: not implemented"; return 0 }

// GetCalls gets the received messages for specific call type
func (s *Greeter) GetCalls(key CallType) [][]*HelloRequest { _ = "STUB: not implemented"; return nil }

// GetSendCounts gets the stream send counts
func (s *Greeter) GetSendCounts(key CallType) map[int]int { _ = "STUB: not implemented"; return nil }

// GetConnectionCount gets the connection count
func (s *Greeter) GetConnectionCount() int { _ = "STUB: not implemented"; return 0 }

// NewGreeter creates new greeter server
func NewGreeter() *Greeter { _ = "STUB: not implemented"; return nil }

// NewHWStats creates new stats handler
func NewHWStats() *HWStatsHandler { _ = "STUB: not implemented"; return nil }

// HWStatsHandler is for gRPC stats
type HWStatsHandler struct {
	mutex     *sync.RWMutex
	connCount int
}

// GetConnectionCount gets the connection count
func (c *HWStatsHandler) GetConnectionCount() int { _ = "STUB: not implemented"; return 0 }

// GetCountByWorker gets count of requests by goroutine
func (s *Greeter) GetCountByWorker(key CallType) map[string]int {
	_ = "STUB: not implemented"
	return nil
}

// HandleConn handle the connection
func (c *HWStatsHandler) HandleConn(ctx context.Context, cs stats.ConnStats) {
	_ = "STUB: not implemented"
	// no-op

	// TagConn exists to satisfy gRPC stats.Handler.
	return
}

func (c *HWStatsHandler) TagConn(ctx context.Context, cti *stats.ConnTagInfo) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// HandleRPC implements per-RPC tracing and stats instrumentation.
func (c *HWStatsHandler) HandleRPC(ctx context.Context, rs stats.RPCStats) {
	_ = "STUB: not implemented"
	// no-op

	// TagRPC implements per-RPC context management.
	return
}

func (c *HWStatsHandler) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
