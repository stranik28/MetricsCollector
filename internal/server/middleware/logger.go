package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/stranik28/MetricsCollector/internal/server/logger"
	"go.uber.org/zap"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		requestURL := c.Request.RequestURI

		c.Next()

		// after request
		latency := time.Since(t)
		// access the status we are sending
		status := c.Writer.Status()
		responseSize := c.Writer.Size()

		// Log request and response information
		logger.Log.Info("Response field", zap.Int("STATUS_CODE", status),
			zap.Int("RESPONSE_SIZE", responseSize))
		logger.Log.Info("Request field", zap.Duration("LATENCY", latency),
			zap.String("REQUEST_URL", requestURL))
	}
}
