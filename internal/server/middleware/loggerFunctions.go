package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

type GinContextWithLogger struct {
	*gin.Context
	*zap.Logger
}

func Init(level string) (*zap.Logger, error) {
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return nil, err
	}

	cfg := zap.NewProductionConfig()

	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
	cfg.Level = lvl
	cfg.Encoding = "json"
	cfg.OutputPaths = []string{"info.log"}

	zl, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	defer zl.Sync()
	// устанавливаем синглтон

	return zl, nil
}
