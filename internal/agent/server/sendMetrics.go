package server

import (
	"fmt"
	"github.com/stranik28/MetricsCollector/internal/agent/logger"
	"github.com/stranik28/MetricsCollector/internal/agent/models"
	"github.com/stranik28/MetricsCollector/internal/agent/storage"
	"go.uber.org/zap"
)

func SendMetrics(memStorage *storage.MemStorage, servAddr string) {
	for _, store := range memStorage.Metrics {
		for k, v := range store.Gauge {
			model := models.Metrics{
				ID:    k,
				MType: "gauge",
				Value: &v,
			}
			url := fmt.Sprintf("http://%s/update/", servAddr)
			req := NewServer(url)
			code := req.SendReqPost("POST", model)
			if code != 200 {
				logger.Log.Error("Failed to send metrics", zap.Int("code", code), zap.String("url", url))
			}
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
			logger.Log.Error("Metrics not sent", zap.Any("Metrics", model))
		}
	}
	memStorage.ClearMemStorage()
}
