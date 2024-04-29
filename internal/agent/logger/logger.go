package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var Log *zap.Logger = zap.NewNop()

func Init(level string) error {
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return err
	}
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
	cfg.Level = lvl
	cfg.Encoding = "json"
	cfg.OutputPaths = []string{"agent.log"}
	zl, err := cfg.Build()
	if err != nil {
		return err
	}

	defer zl.Sync()
	Log = zl
	return nil
}
