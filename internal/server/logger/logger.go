package logger

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var Log *zap.Logger = zap.NewNop()

func MiddlewareInit() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		requestUrl := c.Request.RequestURI
		c.Next()

		// after request
		latency := time.Since(t)
		// access the status we are sending
		status := c.Writer.Status()
		responseSize := c.Writer.Size()

		Log.Info("Response filed", zap.Int("STATUS_CODE", status),
			zap.Int("RESPONSE_SIZE", responseSize))
		Log.Info("Request field", zap.Duration("LATENCY", latency),
			zap.String("REQUEST URL", requestUrl))
	}
}

func Init(level string) error {
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return err
	}
	// создаём новую конфигурацию логера
	cfg := zap.NewProductionConfig()
	// устанавливаем уровень
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
	cfg.Level = lvl
	cfg.Encoding = "json"
	cfg.OutputPaths = []string{"info.log"}
	// создаём логер на основе конфигурации
	zl, err := cfg.Build()
	if err != nil {
		return err
	}

	defer zl.Sync()
	// устанавливаем синглтон
	Log = zl
	return nil
}
