package collector

import (
	"fmt"
	"github.com/stranik28/MetricsCollector/internal/agent/server"
	"github.com/stranik28/MetricsCollector/internal/agent/storage"
	"go.uber.org/zap"
	"time"
)

func MetricsCollector(flagReportInterval int, flagPollInterval int, flagServAddr string, logger *zap.Logger, secretKey string) error {
	count := 0
	memStorage := storage.NewMemStorage()
	for {
		logger.Info("Collecting metrics...")
		metric, err := collectMetrics()
		for key, _ := range metric.Gauge {
			fmt.Println(key)
		}
		if err != nil {
			logger.Error("Error collecting metrics", zap.Error(err))
		}
		memStorage.AddMetric(metric)
		if count >= flagPollInterval {
			count = 0
			logger.Info("Polling metrics...")
			server.SendMetrics(memStorage, flagServAddr, logger, secretKey)
		}
		time.Sleep(time.Duration(flagReportInterval) * time.Second)
		count += flagReportInterval
	}
}
