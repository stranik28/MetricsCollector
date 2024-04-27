package handlers

import (
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
	logger.Log.Debug("Getting JSON")

	var req models.Metrics

	if err := c.Bind(&req); err != nil {
		logger.Log.Debug("cannot decode request JSON body", zap.Error(err))
		c.AbortWithStatus(http.StatusInternalServerError)
	}

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
