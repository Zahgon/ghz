package main

import (
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/alecthomas/kingpin"
	"go.uber.org/zap"

	"github.com/bojand/ghz/printer"
	"github.com/bojand/ghz/runner"
)

var (
	// set by goreleaser with -ldflags="-X main.version=..."
	version = "dev"

	nCPUs = runtime.GOMAXPROCS(-1)

	cPath = kingpin.Flag("config", "Path to the JSON or TOML config file that specifies all the test run settings.").PlaceHolder(" ").String()

	// Proto
	isProtoSet = false
	proto      = kingpin.Flag("proto", `The Protocol Buffer .proto file.`).
			PlaceHolder(" ").IsSetByUser(&isProtoSet).String()

	isProtoSetSet = false
	protoset      = kingpin.Flag("protoset", "The compiled protoset file. Alternative to proto. -proto takes precedence.").
			PlaceHolder(" ").IsSetByUser(&isProtoSetSet).String()

	isCallSet = false
	call      = kingpin.Flag("call", `A fully-qualified method name in 'package.Service/method' or 'package.Service.Method' format.`).
			PlaceHolder(" ").IsSetByUser(&isCallSet).String()

	isImportSet = false
	paths       = kingpin.Flag("import-paths", "Comma separated list of proto import paths. The current working directory and the directory of the protocol buffer file are automatically added to the import list.").
			Short('i').PlaceHolder(" ").IsSetByUser(&isImportSet).String()

	// Security
	isCACertSet = false
	cacert      = kingpin.Flag("cacert", "File containing trusted root certificates for verifying the server.").
			PlaceHolder(" ").IsSetByUser(&isCACertSet).String()

	isCertSet = false
	cert      = kingpin.Flag("cert", "File containing client certificate (public key), to present to the server. Must also provide -key option.").
			PlaceHolder(" ").IsSetByUser(&isCertSet).String()

	isKeySet = false
	key      = kingpin.Flag("key", "File containing client private key, to present to the server. Must also provide -cert option.").
			PlaceHolder(" ").IsSetByUser(&isKeySet).String()

	isCNameSet = false
	cname      = kingpin.Flag("cname", "Server name override when validating TLS certificate - useful for self signed certs.").
			PlaceHolder(" ").IsSetByUser(&isCNameSet).String()

	isSkipSet  = false
	skipVerify = kingpin.Flag("skipTLS", "Skip TLS client verification of the server's certificate chain and host name.").
			Default("false").IsSetByUser(&isSkipSet).Bool()

	isInsecSet = false
	insecure   = kingpin.Flag("insecure", "Use plaintext and insecure connection.").
			Default("false").IsSetByUser(&isInsecSet).Bool()

	isAuthSet = false
	authority = kingpin.Flag("authority", "Value to be used as the :authority pseudo-header. Only works if -insecure is used.").
			PlaceHolder(" ").IsSetByUser(&isAuthSet).String()

	// Run
	isAsyncSet = false
	async      = kingpin.Flag("async", "Make requests asynchronous as soon as possible. Does not wait for request to finish before sending next one.").
			Default("false").IsSetByUser(&isAsyncSet).Bool()

	isRPSSet = false
	rps      = kingpin.Flag("rps", "Requests per second (RPS) rate limit for constant load schedule. Default is no rate limit.").
			Default("0").Short('r').IsSetByUser(&isRPSSet).Uint()

	isScheduleSet = false
	schedule      = kingpin.Flag("load-schedule", "Specifies the load schedule. Options are const, step, or line. Default is const.").
			Default("const").IsSetByUser(&isScheduleSet).String()

	isLoadStartSet = false
	loadStart      = kingpin.Flag("load-start", "Specifies the RPS load start value for step or line schedules.").
			Default("0").IsSetByUser(&isLoadStartSet).Uint()

	isLoadStepSet = false
	loadStep      = kingpin.Flag("load-step", "Specifies the load step value or slope value.").
			Default("0").IsSetByUser(&isLoadStepSet).Int()

	isLoadEndSet = false
	loadEnd      = kingpin.Flag("load-end", "Specifies the load end value for step or line load schedules.").
			Default("0").IsSetByUser(&isLoadEndSet).Uint()

	isLoadStepDurSet = false
	loadStepDuration = kingpin.Flag("load-step-duration", "Specifies the load step duration value for step load schedule.").
				Default("0").IsSetByUser(&isLoadStepDurSet).Duration()

	isLoadMaxDurSet = false
	loadMaxDuration = kingpin.Flag("load-max-duration", "Specifies the max load duration value for step or line load schedule.").
			Default("0").IsSetByUser(&isLoadMaxDurSet).Duration()

	// Concurrency
	isCSet = false
	c      = kingpin.Flag("concurrency", "Number of request workers to run concurrently for const concurrency schedule. Default is 50.").
		Short('c').Default("50").IsSetByUser(&isCSet).Uint()

	isCScheduleSet = false
	cschdule       = kingpin.Flag("concurrency-schedule", "Concurrency change schedule. Options are const, step, or line. Default is const.").
			Default("const").IsSetByUser(&isCScheduleSet).String()

	isCStartSet = false
	cStart      = kingpin.Flag("concurrency-start", "Concurrency start value for step and line concurrency schedules.").
			Default("0").IsSetByUser(&isCStartSet).Uint()

	isCEndSet = false
	cEnd      = kingpin.Flag("concurrency-end", "Concurrency end value for step and line concurrency schedules.").
			Default("0").IsSetByUser(&isCEndSet).Uint()

	isCStepSet = false
	cstep      = kingpin.Flag("concurrency-step", "Concurrency step / slope value for step and line concurrency schedules.").
			Default("1").IsSetByUser(&isCStepSet).Int()

	isCStepDurSet = false
	cStepDuration = kingpin.Flag("concurrency-step-duration", "Specifies the concurrency step duration value for step concurrency schedule.").
			Default("0").IsSetByUser(&isCStepDurSet).Duration()

	isCMaxDurSet = false
	cMaxDuration = kingpin.Flag("concurrency-max-duration", "Specifies the max concurrency adjustment duration value for step or line concurrency schedule.").
			Default("0").IsSetByUser(&isCMaxDurSet).Duration()

	// Other
	isNSet = false
	n      = kingpin.Flag("total", "Number of requests to run. Default is 200.").
		Short('n').Default("200").IsSetByUser(&isNSet).Uint()

	isTSet = false
	t      = kingpin.Flag("timeout", "Timeout for each request. Default is 20s, use 0 for infinite.").
		Default("20s").Short('t').IsSetByUser(&isTSet).Duration()

	isZSet = false
	z      = kingpin.Flag("duration", "Duration of application to send requests. When duration is reached, application stops and exits. If duration is specified, n is ignored. Examples: -z 10s -z 3m.").
		Short('z').Default("0").IsSetByUser(&isZSet).Duration()

	isXSet = false
	x      = kingpin.Flag("max-duration", "Maximum duration of application to send requests with n setting respected. If duration is reached before n requests are completed, application stops and exits. Examples: -x 10s -x 3m.").
		Short('x').Default("0").IsSetByUser(&isXSet).Duration()

	isZStopSet = false
	zstop      = kingpin.Flag("duration-stop", "Specifies how duration stop is reported. Options are close, wait or ignore. Default is close.").
			Default("close").IsSetByUser(&isZStopSet).String()

	// Data
	isDataSet = false
	data      = kingpin.Flag("data", "The call data as stringified JSON. If the value is '@' then the request contents are read from stdin.").
			Short('d').PlaceHolder(" ").IsSetByUser(&isDataSet).String()

	isDataPathSet = false
	dataPath      = kingpin.Flag("data-file", "File path for call data JSON file. Examples: /home/user/file.json or ./file.json.").
			Short('D').PlaceHolder("PATH").PlaceHolder(" ").IsSetByUser(&isDataPathSet).String()

	isBinDataSet = false
	binData      = kingpin.Flag("binary", "The call data comes as serialized binary message or multiple count-prefixed messages read from stdin.").
			Short('b').Default("false").IsSetByUser(&isBinDataSet).Bool()

	isBinDataPathSet = false
	binPath          = kingpin.Flag("binary-file", "File path for the call data as serialized binary message or multiple count-prefixed messages.").
				Short('B').PlaceHolder(" ").IsSetByUser(&isBinDataPathSet).String()

	isMDSet = false
	md      = kingpin.Flag("metadata", "Request metadata as stringified JSON.").
		Short('m').PlaceHolder(" ").IsSetByUser(&isMDSet).String()

	isMDPathSet = false
	mdPath      = kingpin.Flag("metadata-file", "File path for call metadata JSON file. Examples: /home/user/metadata.json or ./metadata.json.").
			Short('M').PlaceHolder(" ").IsSetByUser(&isMDPathSet).String()

	isSISet = false
	si      = kingpin.Flag("stream-interval", "Interval for stream requests between message sends.").
		Default("0").IsSetByUser(&isSISet).Duration()

	isSCSet = false
	scd     = kingpin.Flag("stream-call-duration", "Duration after which client will close the stream in each streaming call.").
		Default("0").IsSetByUser(&isSCSet).Duration()

	isSCCSet = false
	scc      = kingpin.Flag("stream-call-count", "Count of messages sent, after which client will close the stream in each streaming call.").
			Default("0").IsSetByUser(&isSCCSet).Uint()

	isSDMSet = false
	sdm      = kingpin.Flag("stream-dynamic-messages", "In streaming calls, regenerate and apply call template data on every message send.").
			Default("false").IsSetByUser(&isSDMSet).Bool()

	isRMDSet = false
	rmd      = kingpin.Flag("reflect-metadata", "Reflect metadata as stringified JSON used only for reflection request.").
			PlaceHolder(" ").IsSetByUser(&isRMDSet).String()

	// Output
	isOutputSet = false
	output      = kingpin.Flag("output", "Output path. If none provided stdout is used.").
			Short('o').PlaceHolder(" ").IsSetByUser(&isOutputSet).String()

	isFormatSet = false
	format      = kingpin.Flag("format", "Output format. One of: summary, csv, json, pretty, html, influx-summary, influx-details. Default is summary.").
			Short('O').Default("summary").PlaceHolder(" ").IsSetByUser(&isFormatSet).Enum("summary", "csv", "json", "pretty", "html", "influx-summary", "influx-details", "prometheus")

	isSkipFirstSet = false
	skipFirst      = kingpin.Flag("skipFirst", "Skip the first X requests when doing the results tally.").
			Default("0").IsSetByUser(&isSkipFirstSet).Uint()

	isCESet     = false
	countErrors = kingpin.Flag("count-errors", "Count erroneous (non-OK) resoponses in stats calculations.").
			Default("false").IsSetByUser(&isCESet).Bool()

	// Connection
	isConnSet = false
	conns     = kingpin.Flag("connections", "Number of connections to use. Concurrency is distributed evenly among all the connections. Default is 1.").
			Default("1").IsSetByUser(&isConnSet).Uint()

	isCTSet = false
	ct      = kingpin.Flag("connect-timeout", "Connection timeout for the initial connection dial. Default is 10s.").
		Default("10s").IsSetByUser(&isCTSet).Duration()

	isKTSet = false
	kt      = kingpin.Flag("keepalive", "Keepalive time duration. Only used if present and above 0.").
		Default("0").IsSetByUser(&isKTSet).Duration()

	// Meta
	isNameSet = false
	name      = kingpin.Flag("name", "User specified name for the test.").
			PlaceHolder(" ").IsSetByUser(&isNameSet).String()

	isTagsSet = false
	tags      = kingpin.Flag("tags", "JSON representation of user-defined string tags.").
			PlaceHolder(" ").IsSetByUser(&isTagsSet).String()

	isCPUSet = false
	cpus     = kingpin.Flag("cpus", "Number of cpu cores to use.").
			Default(strconv.FormatUint(uint64(nCPUs), 10)).IsSetByUser(&isCPUSet).Uint()

	// Debug
	isDebugSet = false
	debug      = kingpin.Flag("debug", "The path to debug log file.").
			PlaceHolder(" ").IsSetByUser(&isDebugSet).String()

	isEnableCompressionSet = false
	enableCompression      = kingpin.Flag("enable-compression", "Enable Gzip compression on requests.").
				Short('e').Default("false").IsSetByUser(&isEnableCompressionSet).Bool()

	isLBStrategySet = false
	lbStrategy      = kingpin.Flag("lb-strategy", "Client load balancing strategy.").
			PlaceHolder(" ").IsSetByUser(&isLBStrategySet).String()

	// message size
	isMaxRecvMsgSizeSet = false
	maxRecvMsgSize      = kingpin.Flag("max-recv-message-size", "Maximum message size the client can receive.").
				PlaceHolder(" ").IsSetByUser(&isMaxRecvMsgSizeSet).String()

	isMaxSendMsgSizeSet = false
	maxSendMsgSize      = kingpin.Flag("max-send-message-size", "Maximum message size the client can send.").
				PlaceHolder(" ").IsSetByUser(&isMaxSendMsgSizeSet).String()

	isDisableTemplateFuncsSet = false
	disableTemplateFuncs      = kingpin.Flag("disable-template-functions", "Do not use and execute any template functions in call template data. Useful for better performance").
					Default("false").IsSetByUser(&isDisableTemplateFuncsSet).Bool()

	isDisableTemplateDataSet = false
	disableTemplateData      = kingpin.Flag("disable-template-data", "Do not use and execute any call template data. Useful for better performance.").
					Default("false").IsSetByUser(&isDisableTemplateDataSet).Bool()

	// host main argument
	isHostSet = false
	host      = kingpin.Arg("host", "Host and port to test.").String()
)

func main() {
	kingpin.Version(version)
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.CommandLine.VersionFlag.Short('v')
	kingpin.Parse()

	isHostSet = *host != ""

	cfgPath := strings.TrimSpace(*cPath)

	var cfg runner.Config

	if cfgPath != "" {
		err := runner.LoadConfig(cfgPath, &cfg)
		kingpin.FatalIfError(err, "")

		args := os.Args[1:]
		if len(args) > 1 {
			var cmdCfg runner.Config
			err = createConfigFromArgs(&cmdCfg)
			kingpin.FatalIfError(err, "")

			err = mergeConfig(&cfg, &cmdCfg)
			kingpin.FatalIfError(err, "")
		}
	} else {
		err := createConfigFromArgs(&cfg)

		kingpin.FatalIfError(err, "")
	}

	var logger *zap.SugaredLogger

	options := []runner.Option{runner.WithConfig(&cfg)}
	if len(cfg.Debug) > 0 {
		var err error
		logger, err = createLogger(cfg.Debug)
		kingpin.FatalIfError(err, "")

		defer logger.Sync()

		options = append(options, runner.WithLogger(logger))
	}

	if logger != nil && isLBStrategySet && cfg.Host != "" && !strings.HasPrefix(cfg.Host, "dns:///") {
		logger.Warnw("Load balancing strategy set without using DNS (dns:///) scheme", "strategy", cfg.LBStrategy, "host", cfg.Host)
	}

	if logger != nil {
		logger.Debugw("Start Run", "config", cfg)
	}

	report, err := runner.Run(cfg.Call, cfg.Host, options...)
	if err != nil {
		if logger != nil {
			logger.Errorf("Error from run: %+v", err.Error())
		}

		handleError(err)
	}

	output := os.Stdout
	outputPath := strings.TrimSpace(cfg.Output)

	if logger != nil {
		logger.Debug("Run finished")
	}

	if outputPath != "" {
		f, err := os.Create(outputPath)
		if err != nil {
			if logger != nil {
				logger.Errorw("Error opening file "+outputPath+": "+err.Error(),
					"error", err)
			}

			handleError(err)
		}

		defer func() {
			handleError(f.Close())
		}()

		output = f
	}

	if logger != nil {
		logPath := "stdout"
		if outputPath != "" {
			logPath = outputPath
		}

		logger.Debugw("Printing report to "+logPath, "path", logPath)
	}

	p := printer.ReportPrinter{
		Report: report,
		Out:    output,
	}

	handleError(p.Print(cfg.Format))
}

func handleError(err error) { _ = "STUB: not implemented"; return }

func createConfigFromArgs(cfg *runner.Config) error { _ = "STUB: not implemented"; return nil }

func mergeConfig(dest *runner.Config, src *runner.Config) error {
	_ = "STUB: not implemented"
	return nil
}

// proto

// security

// run

// data

// other

// load

// concurrency

// message size

// call data template functions behavior

// call data template behavior

// createLogger creates a new zap logger
func createLogger(path string) (*zap.SugaredLogger, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
