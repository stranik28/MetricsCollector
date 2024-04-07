package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stranik28/MetricsCollector/internal/server/service"
	"github.com/stranik28/MetricsCollector/internal/server/storage"
	"net/http"
)

func GetMetric(c *gin.Context) {
	metricName := c.Param("metricName")
	metricType := c.Param("metricType")
	val, err := service.GetMetricByName(metricName, metricType)
	if err != nil {
		if errors.Is(err, storage.ErrorMetricsNotFound) {
			c.JSON(http.StatusNotFound, err.Error())
		}
	}
	c.JSON(http.StatusOK, val)

}
