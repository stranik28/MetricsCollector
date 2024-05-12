package service

import (
	"github.com/stranik28/MetricsCollector/internal/server/models"
	"github.com/stranik28/MetricsCollector/internal/server/storage"
)

func UpdateMetrics(reqModels []models.Metrics) ([]models.Metrics, error) {
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
			storage.SetMemStorage(reqModel.ID, val)
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
	return respModels, nil
}
