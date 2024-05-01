package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stranik28/MetricsCollector/internal/server/logger"
	"github.com/stranik28/MetricsCollector/internal/server/models"
	"github.com/stranik28/MetricsCollector/internal/server/service"
	"github.com/stranik28/MetricsCollector/internal/server/storage"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
)

func UpdateMetrics(c *gin.Context) {
	var req models.Metrics
	var buf bytes.Buffer
	// читаем тело запроса
	_, err := buf.ReadFrom(c.Request.Body)
	if err != nil {
		logger.Log.Error("Error reading body", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// десериализуем JSON в Visitor
	if err = json.Unmarshal(buf.Bytes(), &req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logger.Log.Debug("received request", zap.Any("request", req))

	responseModel, err := service.UpdateMetrics(req)

	if err != nil {
		if errors.Is(err, storage.ErrorIncorrectTypeMetrics) || errors.Is(err, storage.ErrorIncorrectTypeInt64) ||
			errors.Is(err, storage.ErrorIncorrectTypeFloat64) {
			c.JSON(http.StatusBadRequest, err.Error())
		}
	}

	c.JSON(http.StatusOK, responseModel)
}

func UpdateMetricsParam(c *gin.Context) {
	var req models.Metrics
	metricType := c.Param("metricType")
	metricName := c.Param("metricName")
	metricValue := c.Param("metricValue")

	req.ID = metricName
	req.MType = metricType

	if metricType == "counter" {
		value, err := strconv.ParseInt(strings.TrimSpace(metricValue), 10, 64)
		if err != nil {
			err = storage.ErrorIncorrectTypeInt64
			c.JSON(http.StatusBadRequest, err.Error())
		}
		req.Delta = &value
	} else if metricType == "gauge" {
		value, err := strconv.ParseFloat(strings.TrimSpace(metricValue), 64)
		if err != nil {
			err = storage.ErrorIncorrectTypeFloat64
			c.JSON(http.StatusBadRequest, err.Error())
		}
		req.Value = &value
	}

	_, err := service.UpdateMetrics(req)
	if err != nil {
		if errors.Is(err, storage.ErrorIncorrectTypeMetrics) || errors.Is(err, storage.ErrorIncorrectTypeInt64) ||
			errors.Is(err, storage.ErrorIncorrectTypeFloat64) {
			c.JSON(http.StatusBadRequest, err.Error())
		}
	}
	c.JSON(http.StatusOK, nil)
}
