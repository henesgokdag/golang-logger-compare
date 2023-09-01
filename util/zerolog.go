package util

import (
	"github.com/rs/zerolog"
	zerolog_logger "github.com/rs/zerolog/log"
)

func setZerologLogLevel() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

func loggerZerologWithoutStruct() {
	zerolog_logger.Debug().Msg("this is a debug message")
	zerolog_logger.Info().Msg("this is an info message")
	zerolog_logger.Warn().Msg("this is a warning message")
	zerolog_logger.Error().Msg("this is an error message")
}
