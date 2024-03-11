package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"runtime/debug"
	"strings"
)

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
