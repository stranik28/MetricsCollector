package service

import (
	"context"
	"github.com/stranik28/MetricsCollector/internal/server"
	"github.com/stranik28/MetricsCollector/internal/server/models"
	"github.com/stranik28/MetricsCollector/internal/server/storage"
)

func UpdateMetrics(c context.Context, reqModels []models.Metrics) ([]models.Metrics, error) {
	respModels := make([]models.Metrics, len(reqModels))
	for _, reqModel := range reqModels {
		switch reqModel.MType {
		case "counter":
			val, ok := storage.GetMemStorage(reqModel.ID)
			if !ok {
				val = storage.Metric{
					Gauge:   0,
					Counter: *reqModel.Delta,
				}
			} else {
				val.Counter += *reqModel.Delta
			}
			if server.DBURL == "" {
				storage.SetMemStorage(reqModel.ID, val)
			}
			reqModel.Value = &val.Gauge
		case "gauge":
			val, ok := storage.GetMemStorage(reqModel.ID)
			if !ok {
				val = storage.Metric{
					Gauge:   *reqModel.Value,
					Counter: 0,
				}
			} else {
				val.Gauge = *reqModel.Value
			}
			storage.SetMemStorage(reqModel.ID, val)
			reqModel.Value = &val.Gauge
		default:
			err := storage.ErrorIncorrectTypeMetrics
			return respModels, err
		}
		respModels = append(respModels, reqModel)
	}
	if server.DBURL != "" {
		db, err := storage.Connect(c, server.DBURL)
		if err != nil {
			return nil, err
		}
		defer db.Close()
		if err := storage.InsertMetric(c, db, respModels[1:]); err != nil {
			return nil, err
		}
	}

	return respModels[1:], nil
}
