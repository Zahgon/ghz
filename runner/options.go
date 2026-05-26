package runner

import (
	"io"
	"text/template"
	"time"

	"github.com/bojand/ghz/load"
	"github.com/jhump/protoreflect/desc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// BinaryDataFunc is a function that can be used for provide binary data for request programatically.
// MethodDescriptor of the call is passed to the data function.
// CallData for the request is passed and can be used to access worker id, request number, etc...
type BinaryDataFunc func(mtd *desc.MethodDescriptor, callData *CallData) []byte

// ScheduleConst is a constant load schedule
const ScheduleConst = "const"

// ScheduleStep is the step load schedule
const ScheduleStep = "step"

// ScheduleLine is the line load schedule
const ScheduleLine = "line"

// RunConfig represents the request Configs
type RunConfig struct {
	// call settings
	call               string
	host               string
	proto              string
	importPaths        []string
	protoset           string
	protosetBinary     []byte
	enableCompression  bool
	defaultCallOptions []grpc.CallOption

	// security settings
	creds      credentials.TransportCredentials
	cacert     string
	cert       string
	key        string
	cname      string
	skipVerify bool
	insecure   bool
	authority  string

	// load
	rps              int
	loadStart        uint
	loadEnd          uint
	loadStep         int
	loadSchedule     string
	loadDuration     time.Duration
	loadStepDuration time.Duration

	pacer load.Pacer

	// concurrency
	c             int
	cStart        uint
	cEnd          uint
	cStep         int
	cSchedule     string
	cMaxDuration  time.Duration
	cStepDuration time.Duration

	workerTicker load.WorkerTicker

	// test
	n     int
	async bool

	// number of connections
	nConns int

	// timeouts
	z             time.Duration
	timeout       time.Duration
	dialTimeout   time.Duration
	keepaliveTime time.Duration

	zstop string

	streamInterval        time.Duration
	streamCallDuration    time.Duration
	streamCallCount       uint
	streamDynamicMessages bool

	// lbStrategy
	lbStrategy string

	// TODO consolidate these actual value fields to be implemented via provider funcs
	// data & metadata
	data     []byte
	metadata []byte
	binary   bool

	dataFunc         BinaryDataFunc
	dataProviderFunc DataProviderFunc
	dataStreamFunc   StreamMessageProviderFunc
	mdProviderFunc   MetadataProviderFunc

	funcs template.FuncMap

	// reflection metadata
	rmd map[string]string

	// debug
	hasLog bool
	log    Logger

	// template call data
	disableTemplateFuncs bool
	disableTemplateData  bool

	// misc
	name                          string
	cpus                          int
	tags                          []byte
	skipFirst                     int
	countErrors                   bool
	recvMsgFunc                   StreamRecvMsgInterceptFunc
	streamInterceptorProviderFunc StreamInterceptorProviderFunc
}

// Option controls some aspect of run
type Option func(*RunConfig) error

// NewConfig creates a new RunConfig from the options passed
func NewConfig(call, host string, options ...Option) (*RunConfig, error) {
	_ = "STUB: not implemented"

	// init with defaults
	return nil, nil
}

// apply options

// host and call may have been applied via options
// only override if not present

// fix up durations

// checks

// step value for step schedule or
// slope for line schedule

// step value for step schedule or
// slope for line schedule

// WithConfigFromFile uses a configuration JSON file to populate the RunConfig
//
//	WithConfigFromFile("config.json")
func WithConfigFromFile(file string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithConfigFromReader uses a reader containing JSON data to populate the RunConfig
// See also: WithConfigFromFile
func WithConfigFromReader(reader io.Reader) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithConfig uses the configuration to populate the RunConfig
// See also: WithConfigFromFile, WithConfigFromReader
func WithConfig(cfg *Config) Option { _ = "STUB: not implemented"; return *new(Option) }

// init / fix up durations

// WithCertificate specifies the certificate options for the run
//
//	WithCertificate("client.crt", "client.key")
func WithCertificate(cert, key string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithServerNameOverride specifies the certificate options for the run
func WithServerNameOverride(cname string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAuthority specifies the value to be used as the :authority pseudo-header.
// This only works with WithInsecure option.
func WithAuthority(authority string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRootCertificate specifies the root certificate options for the run
//
//	WithRootCertificate("ca.crt")
func WithRootCertificate(cert string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithInsecure specifies that this run should be done using insecure mode
//
//	WithInsecure(true)
func WithInsecure(insec bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSkipTLSVerify skip client side TLS verification of server certificate
func WithSkipTLSVerify(skip bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTotalRequests specifies the N (number of total requests) setting
//
//	WithTotalRequests(1000)
func WithTotalRequests(n uint) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithConcurrency specifies the C (number of concurrent requests) option
//
//	WithConcurrency(20)
func WithConcurrency(c uint) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRPS specifies the RPS (requests per second) limit option
//
//	WithRPS(10)
func WithRPS(v uint) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRunDuration specifies the Z (total test duration) option
//
//	WithRunDuration(time.Duration(2*time.Minute))
func WithRunDuration(z time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDurationStopAction specifies how run duration (Z) timeout is handled
// Possible options are "close", "ignore", and "wait"
//
//	WithDurationStopAction("ignore")
func WithDurationStopAction(action string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTimeout specifies the timeout for each request
//
//	WithTimeout(time.Duration(20*time.Second))
func WithTimeout(timeout time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDialTimeout specifies the initial connection dial timeout
//
//	WithDialTimeout(time.Duration(20*time.Second))
func WithDialTimeout(dt time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithKeepalive specifies the keepalive timeout
//
//	WithKeepalive(time.Duration(1*time.Minute))
func WithKeepalive(k time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithBinaryData specifies the binary data
//
//	msg := &helloworld.HelloRequest{}
//	msg.Name = "bob"
//	binData, _ := proto.Marshal(msg)
//	WithBinaryData(binData)
func WithBinaryData(data []byte) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithClientLoadBalancing specifies the LB strategy to use
// The strategies has to be self written and pre defined
func WithClientLoadBalancing(strategy string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithBinaryDataFunc specifies the binary data func which will be called on each request
//
//	WithBinaryDataFunc(changeFunc)
func WithBinaryDataFunc(data func(mtd *desc.MethodDescriptor, callData *CallData) []byte) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithBinaryDataFromFile specifies the binary data
//
//	WithBinaryDataFromFile("request_data.bin")
func WithBinaryDataFromFile(path string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDataFromJSON loads JSON data from string
//
//	WithDataFromJSON(`{"name":"bob"}`)
func WithDataFromJSON(data string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithData specifies data as generic data that can be serailized to JSON
func WithData(data interface{}) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDataFromReader loads JSON data from reader
//
//	file, _ := os.Open("data.json")
//	WithDataFromReader(file)
func WithDataFromReader(r io.Reader) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDataFromFile loads JSON data from file
//
//	WithDataFromFile("data.json")
func WithDataFromFile(path string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMetadataFromJSON specifies the metadata to be read from JSON string
//
//	WithMetadataFromJSON(`{"request-id":"123"}`)
func WithMetadataFromJSON(md string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMetadata specifies the metadata to be used as a map
//
//	md := make(map[string]string)
//	md["token"] = "foobar"
//	md["request-id"] = "123"
//	WithMetadata(&md)
func WithMetadata(md map[string]string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMetadataFromFile loads JSON metadata from file
//
//	WithMetadataFromJSON("metadata.json")
func WithMetadataFromFile(path string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithName sets the name of the test run
//
//	WithName("greeter service test")
func WithName(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTags specifies the user defined tags as a map
//
//	tags := make(map[string]string)
//	tags["env"] = "staging"
//	tags["created by"] = "joe developer"
//	WithTags(&tags)
func WithTags(tags map[string]string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCPUs specifies the number of CPU's to be used
//
//	WithCPUs(4)
func WithCPUs(c uint) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSkipFirst is the skipFirst option
func WithSkipFirst(c uint) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCountErrors is the count errors option
func WithCountErrors(v bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithProtoFile specified proto file path and optionally import paths
// We will automatically add the proto file path's directory and the current directory
//
//	WithProtoFile("greeter.proto", []string{"/home/protos"})
func WithProtoFile(proto string, importPaths []string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithProtoset specified protoset file path
//
//	WithProtoset("bundle.protoset")
func WithProtoset(protoset string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithProtosetBinary(b []byte) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithStreamInterval sets the stream interval
func WithStreamInterval(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithStreamCallDuration sets the maximum stream call duration at which point the client will close the stream
func WithStreamCallDuration(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithStreamCallCount sets the stream close count
func WithStreamCallCount(c uint) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithStreamDynamicMessages sets the stream dynamic message generation
func WithStreamDynamicMessages(v bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithReflectionMetadata specifies the metadata to be used as a map
//
//	md := make(map[string]string)
//	md["token"] = "foobar"
//	md["request-id"] = "123"
//	WithReflectionMetadata(&md)
func WithReflectionMetadata(md map[string]string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithConnections specifies the number of gRPC connections to use
//
//	WithConnections(5)
func WithConnections(c uint) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithLogger specifies the logging option
func WithLogger(log Logger) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTemplateFuncs adds additional template functions
func WithTemplateFuncs(funcMap template.FuncMap) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithEnableCompression specifies that requests should be done using gzip Compressor
//
//	WithEnableCompression(true)
func WithEnableCompression(enableCompression bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithLoadSchedule specifies the load schedule
//
//	WithLoadSchedule("const")
func WithLoadSchedule(schedule string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithLoadStart specifies the load start
//
//	WithLoadStart(5)
func WithLoadStart(start uint) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithLoadEnd specifies the load end
//
//	WithLoadEnd(25)
func WithLoadEnd(end uint) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithLoadStep specifies the load step
//
//	WithLoadStep(5)
func WithLoadStep(step int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithLoadStepDuration specifies the load step duration for step schedule
func WithLoadStepDuration(duration time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithLoadDuration specifies the load duration
func WithLoadDuration(duration time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithAsync specifies the async option
func WithAsync(async bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithConcurrencySchedule specifies the concurrency adjustment schedule
//
//	WithConcurrencySchedule("const")
func WithConcurrencySchedule(schedule string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithConcurrencyStart specifies the concurrency start for line or step schedule
//
//	WithConcurrencyStart(5)
func WithConcurrencyStart(v uint) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithConcurrencyEnd specifies the concurrency end value for line or step schedule
//
//	WithConcurrencyEnd(25)
func WithConcurrencyEnd(v uint) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithConcurrencyStep specifies the concurrency step value or slope
//
//	WithConcurrencyStep(5)
func WithConcurrencyStep(step int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithConcurrencyStepDuration specifies the concurrency step duration for step schedule
func WithConcurrencyStepDuration(duration time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithConcurrencyDuration specifies the total concurrency adjustment duration
func WithConcurrencyDuration(duration time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithPacer specified the custom pacer to use
func WithPacer(p load.Pacer) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithWorkerTicker specified the custom worker ticker to use
func WithWorkerTicker(ticker load.WorkerTicker) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithStreamRecvMsgIntercept specified the stream receive intercept function
//
//	WithStreamRecvMsgIntercept(func(msg *dynamic.Message, err error) error {
//		if err == nil && msg != nil {
//			reply := &helloworld.HelloReply{}
//			convertErr := msg.ConvertTo(reply)
//			if convertErr == nil {
//				if reply.GetMessage() == "Hello bar" {
//					return ErrEndStream
//				}
//			}
//		}
//		return nil
//	})
func WithStreamRecvMsgIntercept(fn StreamRecvMsgInterceptFunc) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithStreamInterceptor specifies the stream interceptor provider function
func WithStreamInterceptorProviderFunc(interceptor StreamInterceptorProviderFunc) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithDataProvider provides custom data provider
//
//	WithDataProvider(func(*CallData) ([]*dynamic.Message, error) {
//		protoMsg := &helloworld.HelloRequest{Name: "Bob"}
//		dynamicMsg, err := dynamic.AsDynamicMessage(protoMsg)
//		if err != nil {
//			return nil, err
//		}
//		return []*dynamic.Message{dynamicMsg}, nil
//	}),
func WithDataProvider(fn DataProviderFunc) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMetadataProvider provides custom metadata provider
//
//	WithMetadataProvider(ctd *CallData) (*metadata.MD, error) {
//		return &metadata.MD{"token": []string{"secret"}}, nil
//	}),
func WithMetadataProvider(fn MetadataProviderFunc) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithStreamMessageProvider sets custom stream message provider
//
//	WithStreamMessageProvider(func(cd *CallData) (*dynamic.Message, error) {
//		protoMsg := &helloworld.HelloRequest{Name: cd.WorkerID + ": " + strconv.FormatInt(cd.RequestNumber, 10)}
//		dynamicMsg, err := dynamic.AsDynamicMessage(protoMsg)
//		if err != nil {
//			return nil, err
//		}
//
//		callCounter++
//
//		if callCounter == 5 {
//			err = ErrLastMessage
//		}
//
//		return dynamicMsg, err
//	}),
func WithStreamMessageProvider(fn StreamMessageProviderFunc) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithDefaultCallOptions sets the default CallOptions for calls over the connection.
func WithDefaultCallOptions(opts []grpc.CallOption) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithDisableTemplateFuncs disables template functions in call data
func WithDisableTemplateFuncs(v bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDisableTemplateData disables template data execution in call data
func WithDisableTemplateData(v bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func createClientTransportCredentials(skipVerify bool, cacertFile, clientCertFile, clientKeyFile, cname string) (credentials.TransportCredentials, error) {
	_ = "STUB: not implemented"
	return *new(credentials.TransportCredentials), nil
}

// Load the client certificates from disk

// Create a certificate pool from the certificate authority

// Append the certificates from the CA

func fromConfig(cfg *Config) []Option {
	_ = "STUB: not implemented"
	// set up all the options
	return nil
}

// init / fix up durations

// data

// or binary data
