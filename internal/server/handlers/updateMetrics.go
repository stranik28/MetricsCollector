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
	"strconv"
	"strings"
)

func UpdateMetrics(c *gin.Context) {
	loggerC, ok := c.Get("Logger")
	if !ok {
		c.AbortWithStatusJSON(500, "Logger is not available")
		return
	}
	logger := loggerC.(*zap.Logger)
	var req models.Metrics
	var reqBatch []models.Metrics
	var buf bytes.Buffer

	_, err := buf.ReadFrom(c.Request.Body)
	if err != nil {
		logger.Error("Error reading body", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = json.Unmarshal(buf.Bytes(), &reqBatch); err != nil {
		if err = json.Unmarshal(buf.Bytes(), &req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		reqBatch = make([]models.Metrics, 1)
		reqBatch[0] = req
	}
	logger.Debug("received request", zap.Any("request", req))

	responseModel, err := service.UpdateMetrics(c, reqBatch)

	if err != nil {
		if errors.Is(err, storage.ErrorIncorrectTypeMetrics) || errors.Is(err, storage.ErrorIncorrectTypeInt64) ||
			errors.Is(err, storage.ErrorIncorrectTypeFloat64) {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}
	if len(responseModel) == 1 {
		c.JSON(http.StatusOK, responseModel[0])
		return
	}
	c.JSON(http.StatusOK, responseModel)
}

func UpdateMetricsParam(c *gin.Context) {
	var req models.Metrics
	var reqBatch []models.Metrics
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
			return
		}
		req.Delta = &value
	} else if metricType == "gauge" {
		value, err := strconv.ParseFloat(strings.TrimSpace(metricValue), 64)
		if err != nil {
			err = storage.ErrorIncorrectTypeFloat64
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		req.Value = &value
	}

	reqBatch = make([]models.Metrics, 1)
	reqBatch[0] = req
	_, err := service.UpdateMetrics(c, reqBatch)
	if err != nil {
		if errors.Is(err, storage.ErrorIncorrectTypeMetrics) || errors.Is(err, storage.ErrorIncorrectTypeInt64) ||
			errors.Is(err, storage.ErrorIncorrectTypeFloat64) {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}
	c.JSON(http.StatusOK, nil)
}
