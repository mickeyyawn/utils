package logger

import (
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
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
type ZeroLogOptions struct {
	LogLevel       string // Must be one of: ["trace", "debug", "info", "warn", "error", "critical"]
	LevelFieldName string // LevelFieldName sets the field name for the log level or severity.
	JSON           bool   // JSON enables structured logging output in json (prod) set to false in non prod
	Concise        bool   // this should always be false. you want as much info as possible.
	// Tags are additional fields included at the root level of all logs.
	// These can be useful for example the commit hash of a build, or an environment
	// name like prod/stg/dev
	Tags            map[string]string
	SkipHeaders     []string // SkipHeaders are additional headers which are redacted from the logs
	TimeFieldFormat string
	TimeFieldName   string   // set the name of of the timestamp field
} */

/*
Assume we are in NOT in production...which means we go to the
lowest log level and not in JSON
*/
/*

var DefaultOptions = ZeroLogOptions{
	LogLevel:        "trace",
	LevelFieldName:  "level",
	JSON:            false,
	Concise:         false,
	Tags:            nil,
	SkipHeaders:     nil,
	TimeFieldFormat: time.RFC3339Nano, // standard
	TimeFieldName:   "t",              // maybe change this to ts.. TBD
}

*/

var RUNNING_IN_PRODUCTION = false
var LOG_LEVEL = zerolog.TraceLevel // default to trace. if we are running in prod we will raise to Info

func setGlobalStateForZerolog() {
	zerolog.TimestampFieldName = "t"           // set the name of of the timestamp field to something short
	zerolog.TimeFieldFormat = time.RFC3339Nano // defaults to "time.RFC3339Nano"...explicitly set it so it is clear
	zerolog.LevelFieldName = "l"               // set the name of of the level field to something short
	zerolog.MessageFieldName = "m"             // set the name of of the message field to something short
	zerolog.ErrorFieldName = "err"
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
}

//
// TODO: refactor ^ that variable into a func that defaults to false...
//

// this directly returns an instance of a zero logger. no, we are not wrapping or
// otherwise changing that type.
//
// TODO:  pass in env and region and set those as tags...
// TODO: pass in the build hash of the app ???
func Init(serviceName string, environment string, region string) zerolog.Logger {

	//
	// TODO: make the timestamp fields match polygon... ts ?
	//
	// TODO: add the ability (maybe an enum) to let the caller
	// decide if they want to ouput to console, file, or std.out
	//
	//  https://stackoverflow.com/questions/73730972/zerolog-with-stdout-and-file-logger-adds-additional-message-field-in-the-file
	//

	setGlobalStateForZerolog()

	buildInfo, _ := debug.ReadBuildInfo()
	//service := strings.ToLower(serviceName)

	if strings.Contains(strings.ToUpper(os.Getenv("ENVIRONMENT")), "PROD") {
		RUNNING_IN_PRODUCTION = true
		LOG_LEVEL = zerolog.InfoLevel
		//DefaultOptions.JSON = true
		//DefaultOptions.LogLevel = "info"
	}

	//region := os.Getenv("REGION")
	//if region == "" {
	// TODO: log/alert here that a region was not passed in
	// and this logger will default to local/debug...
	//}
	//logLevel, _ := zerolog.Level(LOG_LEVEL)
	//zerolog.SetGlobalLevel(logLevel)
	zerolog.SetGlobalLevel(LOG_LEVEL)

	//log.Logger = log.With().Str("env", "ENVIORNMENT").Logger()
	//log.Logger = log.With().Str("reg", "REGION").Logger()
	//log.Logger = log.With().Str("service", "SERVICE").Logger()

	/*
		if !opts.JSON {
			log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: DefaultOptions.TimeFieldFormat})
		}


	*/

	//logger := log.With().Str("service", service) // .Str("service", service)

	var loggerInstance zerolog.Logger

	if RUNNING_IN_PRODUCTION {
		loggerInstance = zerolog.New(os.Stdout)
		return loggerInstance.With().Str("reg", region).
			Str("env", environment).
			Str("service", serviceName).
			Str("go_v", buildInfo.GoVersion).Timestamp().Caller().Logger()
	} else {
		//loggerInstance := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Str("service", service).Caller()
		loggerInstance = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})
		return loggerInstance.With().Timestamp().Caller().Logger()
	}

	/*

	 logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	    logger.UpdateContext(func(c zerolog.Context) zerolog.Context {
	        return c.Str("name", "john")
	    })

	*/

	//loggerInstance = loggerInstance.With().Str("go_version", buildInfo.GoVersion)

	/*

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

	*/

	//return loggerInstance.Logger()
	//logger := zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()

	//logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})

	//return logger.Logger()

	// TODO: how to output Region, Environment each time !!!
}

/*

multi output, do this in non prod so we can ship that file for demo purposes


func main() {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}

	multi := zerolog.MultiLevelWriter(consoleWriter, os.Stdout)

	logger := zerolog.New(multi).With().Timestamp().Logger()

	logger.Info().Msg("Hello World!")
}

// Output (Line 1: Console; Line 2: Stdout)
// 12:36PM INF Hello World!
// {"level":"info","time":"2019-11-07T12:36:38+03:00","message":"Hello World!"}


func NewLogger(serviceName string, opts ...Options) zerolog.Logger {
	if len(opts) > 0 {
		Configure(opts[0])
	} else {
		Configure(DefaultOptions)
	}
	logger := log.With().Str("service", strings.ToLower(serviceName))
	if !DefaultOptions.Concise && len(DefaultOptions.Tags) > 0 {
		logger = logger.Fields(map[string]interface{}{
			"tags": DefaultOptions.Tags,
		})
	}
	return logger.Logger()
}


*/
