package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func GetLogger() *zap.Logger {
	if logger != nil {
		return logger
	}
	cfg := zap.Config{
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "m",

			LevelKey:    "l",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "t",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "c",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		OutputPaths:      []string{"/tmp/gramya.log"},
		ErrorOutputPaths: []string{"/tmp/gramya-error.log"},
	}

	zapLogger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return zapLogger
}
