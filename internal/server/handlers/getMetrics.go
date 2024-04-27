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

func GetMetric(c *gin.Context) {
	logger.Log.Debug("Getting JSON")
	var req models.Metrics
	dec := json.NewDecoder(c.Request.Body)

	if err := dec.Decode(&req); err != nil {
		logger.Log.Debug("cannot decode request JSON body", zap.Error(err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	responseModel, err := service.GetMetricByName(req)

	if err != nil {
		if errors.Is(err, storage.ErrorMetricsNotFound) {
			c.JSON(http.StatusNotFound, err.Error())
		}
	}

	c.JSON(http.StatusOK, responseModel)
}
