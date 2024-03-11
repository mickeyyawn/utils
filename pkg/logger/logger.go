package logger

import (
	"github.com/rs/zerolog"
	//"github.com/rs/zerolog/log"
	//"github.com/buger/jsonparser"
	"os"
	"strings"
	"time"

	//"io"
	//"encoding/json"

	"runtime/debug"
	//"github.com/getsentry/raven-go"
	//"github.com/getsentry/sentry-go"
	//"encoding/json"
)

/*

example of a singleton logger:  https://stackoverflow.com/questions/18361750/correct-approach-to-global-logging
hook errors and report to sentry:  https://gist.github.com/asdine/f821abe6189a04250ae61b77a3048bd9

*/

/*
func Test() {
    zerolog.SetGlobalLevel(zerolog.ErrorLevel)
    log.Info().Msg("Info message")
    log.Error().Msg("Error message")
} */

/*
Assume we are in NOT in production...which means we go to the
lowest log level and not in JSON
*/
var DefaultOptions = Options{
	LogLevel:        "trace",
	LevelFieldName:  "level",
	JSON:            false,
	Concise:         false,
	Tags:            nil,
	SkipHeaders:     nil,
	TimeFieldFormat: time.RFC3339Nano,
	TimeFieldName:   "timestamp",
}

var RUNNING_IN_PRODUCTION = false

func Init(serviceName string) zerolog.Logger {

	buildInfo, _ := debug.ReadBuildInfo()

	service := strings.ToLower(serviceName)
	if len(service) == 0 {
		// TODO: log/alert here that a service name was not passed and will
		// not be used...
	}

	env := strings.ToUpper(os.Getenv("ENVIRONMENT"))
	//
	// TODO: capitalize the value in there and shorten it to PROD...
	//
	if (env == "PRODUCTION") || (env == "PROD") {
		RUNNING_IN_PRODUCTION = true
		DefaultOptions.JSON = true
		DefaultOptions.LogLevel = "info"
	} else {
		// TODO: log/alert here that an environment was not passed in
		// and this logger will default to local/debug...
	}

	SENTRY_DSN := os.Getenv("SENTRY_DSN")
	if SENTRY_DSN == "" {
		// TODO: log/alert here that Sentry DSN was not set up and will
		// not be used...
		//return newLogger(lvl, os.Stderr), nil, nil
	}

	region := os.Getenv("REGION")
	if region == "" {
		// TODO: log/alert here that a region was not passed in
		// and this logger will default to local/debug...
	}
	logLevel, _ := zerolog.ParseLevel(strings.ToLower(DefaultOptions.LogLevel))
	zerolog.SetGlobalLevel(logLevel)
	zerolog.LevelFieldName = DefaultOptions.LevelFieldName
	zerolog.TimestampFieldName = DefaultOptions.TimeFieldName
	zerolog.TimeFieldFormat = DefaultOptions.TimeFieldFormat

	/*
		if !opts.JSON {
			log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: DefaultOptions.TimeFieldFormat})
		}


	*/

	//logger := log.With().Str("service", service) // .Str("service", service)

	//var loggerInstance zerolog

	if RUNNING_IN_PRODUCTION {
		return zerolog.New(os.Stdout).With().Timestamp().Caller().
			Str("service", service).
			Str("go_version", buildInfo.GoVersion).
			Str("region", region).
			Str("environment", env).Logger()
	} else {
		//loggerInstance := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Str("service", service).Caller()
		return zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Str("service", service).Caller().Logger()
	}

	//return loggerInstance.Logger()
	//logger := zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()

	//logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})

	//return logger.Logger()

	// TODO: how to output Region, Environment each time !!!
}

type Options struct {
	// LogLevel defines the minimum level of severity that app should log.
	//
	// Must be one of: ["trace", "debug", "info", "warn", "error", "critical"]
	LogLevel string

	// LevelFieldName sets the field name for the log level or severity.
	// Some providers parse and search for different field names.
	LevelFieldName string

	// JSON enables structured logging output in json. Make sure to enable this
	// in production mode so log aggregators can receive data in parsable format.
	//
	// In local development mode, its appropriate to set this value to false to
	// receive pretty output and stacktraces to stdout.
	JSON bool

	// Concise mode includes fewer log details during the request flow. For example
	// excluding details like request content length, user-agent and other details.
	// This is useful if during development your console is too noisy.
	Concise bool

	// Tags are additional fields included at the root level of all logs.
	// These can be useful for example the commit hash of a build, or an environment
	// name like prod/stg/dev
	Tags map[string]string

	// SkipHeaders are additional headers which are redacted from the logs
	SkipHeaders []string

	// TimeFieldFormat defines the time format of the Time field, defaulting to "time.RFC3339Nano" see options at:
	// https://pkg.go.dev/time#pkg-constants
	TimeFieldFormat string

	// TimeFieldName sets the field name for the time field.
	// Some providers parse and search for different field names.
	TimeFieldName string
}
