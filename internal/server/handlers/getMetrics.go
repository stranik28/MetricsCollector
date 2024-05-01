package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stranik28/MetricsCollector/internal/server/models"
	"github.com/stranik28/MetricsCollector/internal/server/service"
	"github.com/stranik28/MetricsCollector/internal/server/storage"
	"go.uber.org/zap"
	"net/http"
)

func GetPostMetric(c *gin.Context) {
	loggerC, ok := c.Get("Logger")
	if !ok {
		c.AbortWithStatusJSON(500, "Logger is not available")
		return
	}
	logger := loggerC.(*zap.Logger)
	var req models.Metrics
	logger.Debug("Getting JSON")
	var buf bytes.Buffer

	_, err := buf.ReadFrom(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = json.Unmarshal(buf.Bytes(), &req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logger.Debug("received request", zap.Any("request", req))

	responseModel, err := service.GetMetricByName(req, logger)

	if err != nil {
		if errors.Is(err, storage.ErrorMetricsNotFound) {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, responseModel)
}

func GetMetric(c *gin.Context) {
	loggerC, ok := c.Get("Logger")
	if !ok {
		c.AbortWithStatusJSON(500, "Logger is not available")
		return
	}
	logger := loggerC.(*zap.Logger)
	metricName := c.Param("metricName")
	metricType := c.Param("metricType")
	if !(metricType == "gauge" || metricType == "counter") {
		err := storage.ErrorMetricsNotFound
		c.JSON(http.StatusNotFound, err)
	}
	var req models.Metrics
	req.MType = metricType
	req.ID = metricName
	val, err := service.GetMetricByName(req, logger)
	if err != nil {
		if errors.Is(err, storage.ErrorMetricsNotFound) {
			c.JSON(http.StatusNotFound, err.Error())
		}
	}
	if metricType == "gauge" {
		c.JSON(http.StatusOK, val.Value)
		return
	}
	c.JSON(http.StatusOK, val.Delta)
}
