package service

import (
	"github.com/stranik28/MetricsCollector/internal/server/logger"
	"github.com/stranik28/MetricsCollector/internal/server/models"
	"github.com/stranik28/MetricsCollector/internal/server/storage"
)

func GetMetricByName(modelReq models.Metrics) (models.Metrics, error) {
	metric, ok := storage.GetMemStorage(modelReq.ID)
	if !ok {
		err := storage.ErrorMetricsNotFound
		logger.Log.Warn("Metric not found" + modelReq.ID)
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
		logger.Log.Warn("Metrics not found in storage" + modelReq.MType)
		return modelReq, err
	}

}
