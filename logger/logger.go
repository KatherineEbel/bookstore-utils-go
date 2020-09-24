package logger

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	envLogLevel  = "LOG_LEVEL"
	envLogOutput = "LOG_OUTPUT"
)

var log Logger

type BookStoreLogger interface {
	Printf(string, ...interface{})
	Print(v ...interface{})
}

type Logger struct {
	log *zap.Logger
}

func (l Logger) Print(v ...interface{}) {
	Info(fmt.Sprintf("%v", v))
}

func (l Logger) Printf(format string, v ...interface{}) {
	if len(v) == 0 {
		Info(format)
		return
	}
	Info(fmt.Sprintf(format, v...))
}

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{getOutput()},
		Level:       zap.NewAtomicLevelAt(getLevel()),
		Encoding:    "console",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "msg",
			LevelKey:     "level",
			TimeKey:      "time",
			EncodeLevel:  zapcore.LowercaseColorLevelEncoder,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	var err error
	if log.log, err = logConfig.Build(); err != nil {
		panic(err)
	}
}

func GetLogger() BookStoreLogger {
	return log
}

func Info(msg string, tags ...zap.Field) {
	log.log.Info(msg, tags...)
	_ = log.log.Sync()
}

func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.log.Error(msg, tags...)
	_ = log.log.Sync()
}

func getLevel() zapcore.Level {
	switch os.Getenv(envLogLevel) {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}

func getOutput() string {
	out := strings.TrimSpace(os.Getenv(envLogOutput))
	if out == "" {
		return "stdout"
	}
	return out
}
