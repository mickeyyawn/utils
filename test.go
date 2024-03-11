package main

import (
	"github.com/mickeyyawn/utils/pkg/logger"
)

func main() {
	//
	// init a zerolog logger with the name of the service that
	// is calling it. for demo purpose we will call ourselves "Service B"
	//
	l := logger.Init("Service B")
	l.Info().Msg("This is my information message.")
	l.Error().Msg("Oh noes...something bad happened resulting in an error.")
	//
	// ok, let's pass in some structured information next to the message field
	//
	l.Info().Str("name", "Mr. turtles all the way down").
		Int("age", 42).
		Bool("registered", true).
		Msg("new customer signed up for our product!")
	//
	// write a debug message. Note: this will be ignored if you have
	// set the "ENVIRONMENT" enviroment variable to PROD or PRODUCTION.
	// If set to PROD/PRODUCTION we set the minimum log level to info or above.
	//
	l.Debug().Msg("This is my debug message. Really important debug information here.")

	//
	// now let's pretend we received an error from a caller...
	//
	//err := errors.New("seems we have an error here")
	//l.Error().Err(err).Msg("attaching the error")

}
