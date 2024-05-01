package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func Init(level string, filename string) (*zap.Logger, error) {
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return nil, err
	}
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
	cfg.Level = lvl
	cfg.Encoding = "json"
	cfg.OutputPaths = []string{filename}
	zl, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	defer zl.Sync()

	return zl, nil
}
