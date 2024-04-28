package collector

import (
	"github.com/stranik28/MetricsCollector/internal/agent/logger"
	"github.com/stranik28/MetricsCollector/internal/agent/server"
	"github.com/stranik28/MetricsCollector/internal/agent/storage"
	"go.uber.org/zap"
	"time"
)

func MetricsCollector(flagReportInterval int, flagPollInterval int, flagServAddr string) error {
	count := 0
	memStorage := storage.NewMemStorage()
	err := logger.Init("debug")
	if err != nil {
		return err
	}
	for {
		logger.Log.Info("Collecting metrics...")
		metric, err := collectMetrics()
		if err != nil {
			logger.Log.Error("Error collecting metrics", zap.Error(err))
		}
		memStorage.AddMetric(metric)
		if count >= flagPollInterval {
			count = 0
			logger.Log.Info("Polling metrics...")
			server.SendMetrics(memStorage, flagServAddr)
		}
		time.Sleep(time.Duration(flagReportInterval) * time.Second)
	}
}
