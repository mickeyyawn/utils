package logger

import (
    "github.com/rs/zerolog/log"
)

func Test() {
    zerolog.SetGlobalLevel(zerolog.ErrorLevel)
    log.Info().Msg("Info message")
    log.Error().Msg("Error message")
}
