package runner

import (
	"bytes"
	htmlTemplate "html/template"
	"math/rand"
	"sync"
	"text/template"
	"text/template/parse"
	"time"

	"github.com/Masterminds/sprig/v3"
	"github.com/jhump/protoreflect/desc"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRandPool = sync.Pool{
	New: func() interface{} {
		return rand.New(rand.NewSource(time.Now().UnixNano()))
	},
}

var sprigFuncMap htmlTemplate.FuncMap = sprig.FuncMap()

// CallData represents contextualized data available for templating
type CallData struct {
	WorkerID           string // unique worker ID
	RequestNumber      int64  // unique incremented request number for each request
	FullyQualifiedName string // fully-qualified name of the method call
	MethodName         string // shorter call method name
	ServiceName        string // the service name
	InputName          string // name of the input message type
	OutputName         string // name of the output message type
	IsClientStreaming  bool   // whether this call is client streaming
	IsServerStreaming  bool   // whether this call is server streaming
	Timestamp          string // timestamp of the call in RFC3339 format
	TimestampUnix      int64  // timestamp of the call as unix time in seconds
	TimestampUnixMilli int64  // timestamp of the call as unix time in milliseconds
	TimestampUnixNano  int64  // timestamp of the call as unix time in nanoseconds
	UUID               string // generated UUIDv4 for each call

	t *template.Template
}

var tmplFuncMap = template.FuncMap{
	"newUUID":      newUUID,
	"randomString": randomString,
	"randomInt":    randomInt,
}

// newCallData returns new CallData
func newCallData(
	mtd *desc.MethodDescriptor,
	workerID string, reqNum int64, withFuncs, withTemplateData bool, funcs template.FuncMap) *CallData {
	_ = "STUB: not implemented"
	return nil
}

// Regenerate generates a new instance of call data from this parent instance
// The dynamic data like timestamps and UUIDs are re-filled
func (td *CallData) Regenerate() *CallData { _ = "STUB: not implemented"; return nil }

func (td *CallData) execute(data string) (*bytes.Buffer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This is hacky.
// See https://golang.org/pkg/text/template/#Template
// The *parse.Tree field is exported only for use by html/template
// and should be treated as unexported by all other clients.
func (td *CallData) hasAction(data string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func hasAction(node parse.Node) bool { _ = "STUB: not implemented"; return false }

// ExecuteData applies the call data's parsed template and data string and returns the resulting buffer
func (td *CallData) ExecuteData(data string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (td *CallData) executeMetadata(metadata string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newUUID() string { _ = "STUB: not implemented"; return "" }

const maxLen = 16
const minLen = 2

func stringWithCharset(length int, charset string) string { _ = "STUB: not implemented"; return "" }

func randomString(length int) string { _ = "STUB: not implemented"; return "" }

func randomInt(min, max int) int { _ = "STUB: not implemented"; return 0 }
