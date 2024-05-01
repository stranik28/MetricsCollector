package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/stranik28/MetricsCollector/internal/server/service"
	"go.uber.org/zap"
	"net/http"
)

func AllRecordsHandler(c *gin.Context) {
	loggerC, ok := c.Get("Logger")
	if !ok {
		c.AbortWithStatusJSON(500, "Logger is not available")
		return
	}
	logger := loggerC.(*zap.Logger)
	logger.Info("Update metrics param", zap.Any("params", c.Params))
	metrics, err := service.GetAllMetrics()
	if err != nil {
		logger.Info("Error getting metrics: %v", zap.Error(err))
	}
	c.JSON(http.StatusOK, metrics)
}
