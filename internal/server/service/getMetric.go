package service

import (
	"context"
	"github.com/stranik28/MetricsCollector/internal/server"
	"github.com/stranik28/MetricsCollector/internal/server/models"
	"github.com/stranik28/MetricsCollector/internal/server/storage"
	"go.uber.org/zap"
)

func GetMetricByName(c context.Context, modelReq models.Metrics, logger *zap.Logger) (models.Metrics, error) {
	var metric storage.Metric
	if server.DBURL != "" {
		db, err := storage.NewDBConnection(c, server.DBURL)
		if err != nil {
			logger.Error("Error connecting to database", zap.Error(err))
			return models.Metrics{}, err
		}
		defer db.Conn.Close()
		metric, err = storage.GetMetricByName(c, db, modelReq.ID, modelReq.MType)
		if err != nil {
			logger.Error("Error getting metric from storage", zap.Error(err))
			return models.Metrics{}, err
		}
	} else {
		var ok bool
		metric, ok = storage.GetMemStorage(modelReq.ID)
		if !ok {
			err := storage.ErrorMetricsNotFound
			logger.Warn("Metric not found" + modelReq.ID)
			return modelReq, err
		}
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
