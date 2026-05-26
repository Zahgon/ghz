package protodesc

import (
	"errors"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/grpcreflect"
)

var errNoMethodNameSpecified = errors.New("no method name specified")

// GetMethodDescFromProto gets method descriptor for the given call symbol from proto file given my path proto
// imports is used for import paths in parsing the proto file
func GetMethodDescFromProto(call, proto string, imports []string) (*desc.MethodDescriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetMethodDescFromProtoSet gets method descriptor for the given call symbol from protoset file given my path protoset
func GetMethodDescFromProtoSet(call, protoset string) (*desc.MethodDescriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetMethodDescFromProtoSetBinary gets method descriptor for the given call symbol from protoset binary
func GetMethodDescFromProtoSetBinary(call string, b []byte) (*desc.MethodDescriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetMethodDescFromReflect gets method descriptor for the call from reflection using client
func GetMethodDescFromReflect(call string, client *grpcreflect.Client) (*desc.MethodDescriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getMethodDesc(call string, files map[string]*desc.FileDescriptor) (*desc.MethodDescriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func resolveFileDescriptor(unresolved map[string]*descriptor.FileDescriptorProto, resolved map[string]*desc.FileDescriptor, filename string) (*desc.FileDescriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func findServiceSymbol(resolved map[string]*desc.FileDescriptor, fullyQualifiedName string) (desc.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(desc.Descriptor), nil
}

// parseServiceMethod parses the fully-qualified service name without a leading "."
// and the method name from the input string.
//
// valid inputs:
//
//	package.Service.Method
//	.package.Service.Method
//	package.Service/Method
//	.package.Service/Method
func parseServiceMethod(svcAndMethod string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func newInvalidMethodNameError(svcAndMethod string) error { _ = "STUB: not implemented"; return nil }

func reflectionSupport(err error) error { _ = "STUB: not implemented"; return nil }
