package main

import (
	"github.com/andboson/api-template/pkg/logger"
	
	"github.com/rs/zerolog"
)

func main() {
	l := logger.NewLogger(zerolog.InfoLevel)
	l.Info().Msg("app started")
}
