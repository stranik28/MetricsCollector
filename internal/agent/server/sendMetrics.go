package server

import (
	"fmt"
	"github.com/stranik28/MetricsCollector/internal/agent/logger"
	"github.com/stranik28/MetricsCollector/internal/agent/models"
	"github.com/stranik28/MetricsCollector/internal/agent/storage"
	"go.uber.org/zap"
)

func SendMetrics(memStorage *storage.MemStorage, servAddr string) {
	logger.Log.Info("Memstorage", zap.Any("MemStorage", memStorage))
	for _, store := range memStorage.Metrics {
		logger.Log.Info("Gauge", zap.Any("Gauge", store.Gauge))
		for k, v := range store.Gauge {
			logger.Log.Info("Sending gauge", zap.String("key", k), zap.Any("value", v))
			model := models.Metrics{
				ID:    k,
				MType: "gauge",
				Value: &v,
			}
			logger.Log.Info("Metrics sending", zap.Any("Metrics", model))
			url := fmt.Sprintf("http://%s/update/", servAddr)
			req := NewServer(url)
			code := req.SendReqPost("POST", model)
			logger.Log.Info("Metrics sent", zap.Any("Code", code))
		}
		model := models.Metrics{
			ID:    "PollCount",
			MType: "counter",
			Delta: &store.Counter,
		}
		url := fmt.Sprintf("http://%s/update/", servAddr)
		req := NewServer(url)
		code := req.SendReqPost("POST", model)
		if code != 200 {
			logger.Log.Info("Metrics not sent", zap.Any("Metrics", model))
		}
		logger.Log.Info("Metrics sent", zap.Any("Code", code))
	}
	memStorage.ClearMemStorage()
}
