package logger

import (
	"io"
	"path/filepath"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/rs/zerolog"
)

var logDefaultHeader = map[string]string{
	"time":   "${time_rfc3339_nano}",
	"level":  "${level}",
	"prefix": "${prefix}",
	"file":   "${file}",
	"line":   "${line}",
}

func init() {
	zerolog.CallerMarshalFunc = func(file string, line int) string {
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}
}

type Logger struct {
	*log.Logger
	ZeroLog zerolog.Logger
}

var _ echo.Logger = &Logger{}

func New(writer io.Writer) *Logger {
	l := &Logger{
		Logger:  log.New("-"),
		ZeroLog: zerolog.New(writer).With().Caller().Timestamp().Logger(),
	}

	// log 默认是 ERROR，将 Level 默认都改为 INFO
	l.SetLevel(log.INFO)

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
