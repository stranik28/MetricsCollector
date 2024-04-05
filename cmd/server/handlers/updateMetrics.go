package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stranik28/MetricsCollector/cmd/server/service"
	"github.com/stranik28/MetricsCollector/cmd/server/storage"
	"net/http"
)

func UpdateMetrics(c *gin.Context) {
	metricType := c.Param("metricType")
	metricName := c.Param("metricName")
	metricValue := c.Param("metricValue")
	err := service.UpdateMetrics(metricType, metricValue, metricName)
	if err != nil {
		if errors.Is(err, storage.ErrorIncorrectTypeMetrics) || errors.Is(err, storage.ErrorIncorrectTypeInt64) ||
			errors.Is(err, storage.ErrorIncorrectTypeFloat64) {
			c.JSON(http.StatusBadRequest, err.Error())
		}
	}
	c.JSON(http.StatusOK, nil)
}
