package runner

import (
	"context"
	"time"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/jhump/protoreflect/dynamic/grpcdynamic"
	"google.golang.org/grpc/metadata"
)

// TickValue is the tick value
type TickValue struct {
	instant   time.Time
	reqNumber uint64
}

// Worker is used for doing a single stream of requests in parallel
type Worker struct {
	stub grpcdynamic.Stub
	mtd  *desc.MethodDescriptor

	config   *RunConfig
	workerID string
	active   bool
	stopCh   chan bool
	ticks    <-chan TickValue

	dataProvider     DataProviderFunc
	metadataProvider MetadataProviderFunc
	msgProvider      StreamMessageProviderFunc

	streamRecv                    StreamRecvMsgInterceptFunc
	streamInterceptorProviderFunc StreamInterceptorProviderFunc
}

func (w *Worker) runWorker() error { _ = "STUB: not implemented"; return nil }

// Stop stops the worker. It has to be started with Run() again.
func (w *Worker) Stop() { _ = "STUB: not implemented"; return }

func (w *Worker) makeRequest(tv TickValue) error { _ = "STUB: not implemented"; return nil }

// include the metadata

// RPC errors are handled via stats handler

func (w *Worker) makeUnaryRequest(ctx *context.Context, reqMD *metadata.MD, input *dynamic.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *Worker) makeClientStreamingRequest(ctx *context.Context,
	ctd *CallData, messageProvider StreamMessageProviderFunc) error {
	_ = "STUB: not implemented"
	return nil
}

// default message provider checks counter
// but we also need to keep our own counts
// in case of custom client providers

func (w *Worker) makeServerStreamingRequest(ctx *context.Context, input *dynamic.Message, streamInterceptor StreamInterceptor) error {
	_ = "STUB: not implemented"
	return nil
}

// we should check before receiving a message too

// with any of the cancellation operations we can't just bail
// we have to drain the messages until the server gets the cancel and ends their side of the stream

func (w *Worker) makeBidiRequest(ctx *context.Context,
	ctd *CallData, messageProvider StreamMessageProviderFunc, streamInterceptor StreamInterceptor) error {
	_ = "STUB: not implemented"
	return nil
}

// check at start before send too

// default message provider checks counter
// but we also need to keep our own counts
// in case of custom client providers
