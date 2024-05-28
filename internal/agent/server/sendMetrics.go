package server

import (
	"fmt"
	"github.com/stranik28/MetricsCollector/internal/agent/models"
	"github.com/stranik28/MetricsCollector/internal/agent/storage"
	"go.uber.org/zap"
)

func SendMetrics(metricsChanel chan storage.Metric, servAddr string, logger *zap.Logger, secretKey string) {
	metrics := make([]models.Metrics, 0)
	metricFromChanel := <-metricsChanel
	for k, v := range metricFromChanel.Gauge {
		model := models.Metrics{
			ID:    k,
			MType: "gauge",
			Value: &v,
		}
		metrics = append(metrics, model)
	}
	model := models.Metrics{
		ID:    "PollCount",
		MType: "counter",
		Delta: &metricFromChanel.Counter,
	}
	metrics = append(metrics, model)
	url := fmt.Sprintf("http://%s/update/", servAddr)
	req := NewServer(url)
	code, err := req.SendReqPost("POST", metrics, logger, secretKey)
	if code != 200 || err != nil {
		logger.Error("Failed to send metrics", zap.Int("code", code),
			zap.String("url", url), zap.Error(err))
	}
}
