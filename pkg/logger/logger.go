package logger

import (
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	*zerolog.Logger
}

func NewLogger(level zerolog.Level) *Logger {
	mw := zerolog.MultiLevelWriter(os.Stderr)

	logger := zerolog.New(mw).Level(level).With().Timestamp().Logger()
	return &Logger{&logger}
}
