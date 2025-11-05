package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)


var GlobalLogger *zap.Logger

func init() {
	//logger configuration
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // colors for console
	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // standard time format

	//build the logger instance
	logger, err := cfg.Build()
	if err != nil {
		panic("failed to initialize zap logger: " + err.Error())
	}

	//set global instance
	GlobalLogger = logger
}