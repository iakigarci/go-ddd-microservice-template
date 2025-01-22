package clients

import (
	"sync"

	"github.com/iakigarci/go-ddd-microservice-template/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	instance *zap.Logger
	once     sync.Once
)

func GetLogger(currentLevel config.LogLevel, subsystem string) *zap.Logger {
	once.Do(func() {
		instance = initLogger(currentLevel, subsystem)
	})
	return instance
}

func initLogger(currentLevel config.LogLevel, subsystem string) *zap.Logger {
	var level zapcore.Level
	if currentLevel == config.None {
		return zap.NewNop()
	}

	err := level.UnmarshalText([]byte(currentLevel))
	if err != nil {
		level = zapcore.InfoLevel
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(level),
		Development:      false,
		Encoding:         "json",
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := config.Build(
		zap.AddCallerSkip(1),
		zap.Fields(zap.String("subsystem", subsystem)),
	)
	if err != nil {
		logger, _ = zap.NewProduction()
	}

	return logger
}
