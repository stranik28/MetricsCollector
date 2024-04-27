package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stranik28/MetricsCollector/internal/server/logger"
	"github.com/stranik28/MetricsCollector/internal/server/models"
	"github.com/stranik28/MetricsCollector/internal/server/service"
	"github.com/stranik28/MetricsCollector/internal/server/storage"
	"go.uber.org/zap"
	"net/http"
)

func GetPostMetric(c *gin.Context) {
	var req models.Metrics
	metricName := c.Param("metricName")
	metricType := c.Param("metricType")
	if metricName != "" && metricType != "" {
		req.MType = metricType
		req.ID = metricName
	} else {
		dec := json.NewDecoder(c.Request.Body)

		if err := dec.Decode(&req); err != nil {
			logger.Log.Debug("cannot decode request JSON body", zap.Error(err))
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
	logger.Log.Debug("Getting JSON")

	responseModel, err := service.GetMetricByName(req)

	if err != nil {
		if errors.Is(err, storage.ErrorMetricsNotFound) {
			c.JSON(http.StatusNotFound, err.Error())
		}
	}

	c.JSON(http.StatusOK, responseModel)
}

func GetMetric(c *gin.Context) {
	metricName := c.Param("metricName")
	metricType := c.Param("metricType")
	if !(metricType == "gauge" || metricType == "counter") {
		err := storage.ErrorMetricsNotFound
		c.JSON(http.StatusNotFound, err)
	}
	var req models.Metrics
	req.MType = metricType
	req.ID = metricName
	val, err := service.GetMetricByName(req)
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
