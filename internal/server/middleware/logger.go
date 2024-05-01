package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func Logger(log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		requestURL := c.Request.RequestURI

		c.Set("Logger", log)

		c.Next()

		latency := time.Since(t)
		status := c.Writer.Status()
		responseSize := c.Writer.Size()

		log.Info("Response field", zap.Int("STATUS_CODE", status),
			zap.Int("RESPONSE_SIZE", responseSize))
		log.Info("Request field", zap.Duration("LATENCY", latency),
			zap.String("REQUEST_URL", requestURL))
	}
}
