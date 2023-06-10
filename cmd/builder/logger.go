package builder

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

func NewLogger() *zap.SugaredLogger {
	config := zap.NewProductionConfig()
	config.Encoding = "console"
	config.DisableCaller = true
	config.DisableStacktrace = true
	config.EncoderConfig.TimeKey = "time"
	config.EncoderConfig.LevelKey = "level"
	config.EncoderConfig.MessageKey = "msg"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncoderConfig.ConsoleSeparator = "  "

	logger, err := config.Build()
	if err != nil {
		log.Fatalf("Error creating logger: %v", err)
	}
	return logger.Sugar()
}

func Sync(log *zap.SugaredLogger) {
	_ = log.Sync()
}

func LogLevel(config Configuration) zapcore.Level {
	var level zapcore.Level
	err := json.Unmarshal([]byte(fmt.Sprint(config.LogLevel)), &level)
	if err != nil {
		return zapcore.InfoLevel
	}
	return level
}
