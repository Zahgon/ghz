package runner

import (
	"errors"
	"sync"
	"text/template"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	"google.golang.org/grpc/metadata"
)

// TODO move to own pacakge?
// TODO add tests
// TODO expose public API utilizing only proto API and not dynamic

// ErrEndStream is a signal from message providers that worker should close the stream
// It should not be used for erronous states
var ErrEndStream = errors.New("ending stream")

// ErrLastMessage is a signal from message providers that the returned payload is the last one of the stream
// This is optional but encouraged for optimized performance
// Message payload returned along with this error must be valid and may not be nil
var ErrLastMessage = errors.New("last message")

// DataProviderFunc is the interface for providing data for calls
// For unary and server streaming calls it should return an array with a single element
// For client and bidi streaming calls it should return an array of messages to be used
type DataProviderFunc func(*CallData) ([]*dynamic.Message, error)

// MetadataProviderFunc is the interface for providing metadadata for calls
type MetadataProviderFunc func(*CallData) (*metadata.MD, error)

// StreamMessageProviderFunc is the interface for providing a message for every message send in the course of a streaming call
type StreamMessageProviderFunc func(*CallData) (*dynamic.Message, error)

// StreamRecvMsgInterceptFunc is an interface for function invoked when we receive a stream message
// Clients can return ErrEndStream to end the call early
type StreamRecvMsgInterceptFunc func(*dynamic.Message, error) error

// StreamInterceptorProviderFunc is an interface for a function invoked to generate a stream interceptor
type StreamInterceptorProviderFunc func(*CallData) StreamInterceptor

// StreamInterceptor is an interface for sending and receiving stream messages.
// The interceptor can keep shared state for the send and receive calls.
type StreamInterceptor interface {
	Recv(*dynamic.Message, error) error
	Send(*CallData) (*dynamic.Message, error)
}

type dataProvider struct {
	binary   bool
	data     []byte
	mtd      *desc.MethodDescriptor
	dataFunc BinaryDataFunc

	arrayJSONData []string
	hasActions    bool

	// cached messages only for binary
	mutex          sync.RWMutex
	cachedMessages []*dynamic.Message
}

type mdProvider struct {
	metadata []byte
	preseed  metadata.MD
}

func newDataProvider(mtd *desc.MethodDescriptor,
	binary bool, dataFunc BinaryDataFunc, data []byte,
	withFuncs, withTemplateData bool, funcs template.FuncMap) (*dataProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// fill in JSON string array data for optimization for non client-streaming

// it's an array

// Test if we can preseed data

func (dp *dataProvider) getDataForCall(ctd *CallData) ([]*dynamic.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// try the optimized path for JSON data for non client-streaming

// we want to start from inputs[0] so dec reqNum

func (dp *dataProvider) getMessages(ctd *CallData, i int, inputData []byte) ([]*dynamic.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// only cache JSON data if there are no template actions

// We only cache in case we don't dynamically change the binary message

func newMetadataProvider(mtd *desc.MethodDescriptor, mdData []byte, withFuncs, withTemplateData bool, funcs template.FuncMap) (*mdProvider, error) {
	_ = "STUB: not implemented"
	// Test if we can preseed data
	return nil, nil
}

func (dp *mdProvider) getMetadataForCall(ctd *CallData) (*metadata.MD, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// creates a message from a map
// marshal to JSON then use jsonpb to marshal to message
// this way we follow protobuf more closely and allow camelCase properties.
func messageFromMap(input *dynamic.Message, data *map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func createPayloadsFromJSON(data string, mtd *desc.MethodDescriptor) ([]*dynamic.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createPayloadsFromBinSingleMessage(binData []byte, mtd *desc.MethodDescriptor) ([]*dynamic.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// return empty array if no data

// try to unmarshal input as a single message

func createPayloadsFromBinCountDelimited(binData []byte, mtd *desc.MethodDescriptor) ([]*dynamic.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// return empty array if no data

// try to unmarshal input as several count-delimited messages

func createPayloadsFromBin(binData []byte, mtd *desc.MethodDescriptor) ([]*dynamic.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type dynamicMessageProvider struct {
	mtd           *desc.MethodDescriptor
	data          []byte
	arrayJSONData []string
	arrayLen      uint

	streamCallCount uint
	counter         uint
	indexCounter    uint
}

func newDynamicMessageProvider(mtd *desc.MethodDescriptor, data []byte, streamCallCount uint, withFuncs, withTemplateData bool) (*dynamicMessageProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// it's an array

// Test if we have actions

func (m *dynamicMessageProvider) GetStreamMessage(parentCallData *CallData) (*dynamic.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type staticMessageProvider struct {
	inputs          []*dynamic.Message
	inputLen        uint
	streamCallCount uint
	counter         uint
	indexCounter    uint
}

func newStaticMessageProvider(streamCallCount uint, inputs []*dynamic.Message) (*staticMessageProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *staticMessageProvider) GetStreamMessage(parentCallData *CallData) (*dynamic.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
