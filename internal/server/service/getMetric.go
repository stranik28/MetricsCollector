package service

import (
	"github.com/stranik28/MetricsCollector/internal/server/models"
	"github.com/stranik28/MetricsCollector/internal/server/storage"
	"go.uber.org/zap"
)

func GetMetricByName(modelReq models.Metrics, logger *zap.Logger) (models.Metrics, error) {
	metric, ok := storage.GetMemStorage(modelReq.ID)
	if !ok {
		err := storage.ErrorMetricsNotFound
		logger.Warn("Metric not found" + modelReq.ID)
		return modelReq, err
	}
	if modelReq.MType == "gauge" {
		modelReq.Value = &metric.Gauge
		return modelReq, nil
	} else if modelReq.MType == "counter" {
		modelReq.Delta = &metric.Counter
		return modelReq, nil
	} else {
		err := storage.ErrorMetricsNotFound
		logger.Warn("Metrics not found in storage" + modelReq.MType)
		return modelReq, err
	}

}
