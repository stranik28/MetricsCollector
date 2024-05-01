package collector

import (
	"github.com/stranik28/MetricsCollector/internal/agent/server"
	"github.com/stranik28/MetricsCollector/internal/agent/storage"
	"go.uber.org/zap"
	"time"
)

func MetricsCollector(flagReportInterval int, flagPollInterval int, flagServAddr string, logger *zap.Logger) error {
	count := 0
	memStorage := storage.NewMemStorage()
	for {
		logger.Info("Collecting metrics...")
		metric, err := collectMetrics()
		if err != nil {
			logger.Error("Error collecting metrics", zap.Error(err))
		}
		memStorage.AddMetric(metric)
		if count >= flagPollInterval {
			count = 0
			logger.Info("Polling metrics...")
			server.SendMetrics(memStorage, flagServAddr, logger)
		}
		time.Sleep(time.Duration(flagReportInterval) * time.Second)
		count += flagReportInterval
	}
}
