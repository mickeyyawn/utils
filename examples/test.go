package main

import (
	"context"

	"github.com/mickeyyawn/utils/pkg/logger"
	"github.com/rs/zerolog"
)

var l zerolog.Logger // our global logger for these examples

func main() {
	//
	// init a zerolog logger with the name of the service that
	// is calling it. for demo purpose we will call ourselves "Service 42"
	//
	l = logger.Init("Service 42", "my env", "my reg")
	theBasics()
	decorateLogWithContextualFields()
	usingContext()

	//
	// now let's pretend we received an error from a caller...
	//
	//err := errors.New("seems we have an error here")
	//l.Error().Err(err).Msg("attaching the error")

}

// Just exercise the basics of the logger
func theBasics() {
	l.Info().Msg("This is my information message.")
	l.Error().Msg("Oh noes...something bad happened resulting in an error.")
	//
	// write a debug message. Note: this will be ignored if you have
	// set the "ENVIRONMENT" enviroment variable to PROD or PRODUCTION.
	// If set to PROD/PRODUCTION we set the minimum log level to info or above.
	//
	l.Debug().Msg("This is my debug message. Really important debug information here.")
}

func decorateLogWithContextualFields() {
	//
	// ok, let's pass in some structured information next to the message field
	//
	l.Info().Str("name", "Mr. turtles all the way down").
		Int("age", 42).
		Bool("registered", true).
		Msg("new customer signed up for our product!")
}

func usingContext() {
	ctx := context.Background()
	// Attach the Logger to the context.Context
	ctx = l.WithContext(ctx)
	someFunc(ctx)
}

func someFunc(ctx context.Context) {
	// Get Logger from the go Context. if it's nil, then
	// `zerolog.DefaultContextLogger` is returned, if
	// `DefaultContextLogger` is nil, then a disabled logger is returned.
	//
	//
	// TODO: the following is not correct
	logger := zerolog.Ctx(ctx)
	logger.Info().Msg("Hello")
}
