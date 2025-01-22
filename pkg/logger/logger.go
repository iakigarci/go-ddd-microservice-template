package logger

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

func New(cfg *config.Config) *zap.Logger {
	once.Do(func() {
		instance = initLogger(cfg)
	})
	return instance
}

func initLogger(cfg *config.Config) *zap.Logger {
	var level zapcore.Level
	if cfg.Logging.Level == config.None {
		return zap.NewNop()
	}

	err := level.UnmarshalText([]byte(cfg.Logging.Level))
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
	)
	if err != nil {
		logger, _ = zap.NewProduction()
	}

	return logger
}
