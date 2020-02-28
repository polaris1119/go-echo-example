package logger

import (
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/rs/zerolog"
)

type Logger struct {
	*log.Logger
	ZeroLog zerolog.Logger
}

var _ echo.Logger = &Logger{}

func New(writer io.Writer) *Logger {
	l := &Logger{
		Logger:  log.New("-"),
		ZeroLog: zerolog.New(writer).With().Timestamp().Logger(),
	}

	l.Logger.SetOutput(writer)

	return l
}

func (l *Logger) SetOutput(writer io.Writer) {
	l.Logger.SetOutput(writer)
	l.ZeroLog.Output(writer)
}

func (l *Logger) SetLevel(level log.Lvl) {
	l.Logger.SetLevel(level)
	if level == log.OFF {
		l.ZeroLog = l.ZeroLog.Level(zerolog.Disabled)
	} else {
		zeroLevel := int8(level) - 1
		l.ZeroLog = l.ZeroLog.Level(zerolog.Level(zeroLevel))
	}
}
