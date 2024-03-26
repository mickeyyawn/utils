	err := jsonparser.ObjectEach(*e.buf, func(key, value []byte, vt jsonparser.ValueType, offset int) error {

		//var temp

		switch string(key) {
		case zerolog.MessageFieldName:
			temp := bytesToStrUnsafe(value)
		case zerolog.ErrorFieldName:
			/*
				event.Exception = append(event.Exception, sentry.Exception{
					Value:      bytesToStrUnsafe(value),
					Stacktrace: newStacktrace(),
				})*/
		case zerolog.LevelFieldName, zerolog.TimestampFieldName:
		default:
			//event.Extra[string(key)] = bytesToStrUnsafe(value)
		}

		//return nil
	})

	if err != nil {
		return
	}

	ltwo := logger.InitV2("Service B")
	ltwo.Info("sweeeeet...zap is working", zap.String("username", "johndoe"))







	package logger

import (
		"go.uber.org/zap"
		"go.uber.org/zap/zapcore"
		"log"
		"runtime/debug"
		"strings"
)

type SentryHook struct{}
var SENTRY_DSN = ""
var wg sync.WaitGroup



func InitV2(serviceName string) *zap.Logger {

		buildInfo, _ := debug.ReadBuildInfo()
		goVersion := buildInfo.GoVersion
		log.Println(goVersion)
		service := strings.ToLower(serviceName)
		if len(service) == 0 {
			// TODO: log/alert here that a service name was not passed and will
			// not be used...
		}
		logger := zap.Must(zap.NewProduction())
		logger = logger.WithOptions(zap.Hooks(func(entry zapcore.Entry) error {
			log.Println("well ain't that cool...got a hook working")
			log.Println("and this is the message that was in the entry:")
			log.Println(entry.Message)
			return nil
		}))

		defer logger.Sync()
		return logger
}

	wg.Wait()

if RUNNING_IN_PRODUCTION {
		return zerolog.New(os.Stdout).With().Timestamp().Caller().
			Str("service", service).
			Str("go_version", buildInfo.GoVersion).
			Str("region", region).
			Str("environment", env).Logger().Hook(&SentryHook{})
	} else {
		//loggerInstance := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Str("service", service).Caller()
		return zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Str("service", service).Caller().Logger().Hook(&SentryHook{})
	}



	SENTRY_DSN := os.Getenv("SENTRY_DSN")
	if SENTRY_DSN == "" {
		// TODO: log/alert here that Sentry DSN was not set up and will
		// not be used...
		//return newLogger(lvl, os.Stderr), nil, nil
	}


func (t *SentryHook) Run(
	e *zerolog.Event,
	level zerolog.Level,
	message string,
) {
	log.Println("running the sentry hook!")
	log.Println(message)
	log.Println(level)
	//log.Println(e.)

	//	dec := json.NewDecoder(bytes.NewReader(e.buf))

	/*

		pr, pw := io.Pipe()

		dec := json.NewDecoder(pr)

		err := dec.Decode(&e)
		if err == io.EOF {
			return
		}
		log.Println(e.Time)

	*/

	if level >= zerolog.ErrorLevel {
		wg.Add(1)
		go func() {
			_ = sendToSentry("", message)
			wg.Done()
		}()
	}
}



func sendToSentry(title, msg string) error {

	//
	// TODO: remove this once you make sure this code is working.
	//
	log.Println("sending this Error message to Sentry!!!! woo hoo!!!")

	return nil
}
