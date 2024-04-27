package collector

import (
	"github.com/stranik28/MetricsCollector/internal/agent/server"
	"github.com/stranik28/MetricsCollector/internal/agent/storage"
	"time"
)

func MetricsCollector(flagReportInterval int, flagPollInterval int, flagServAddr string) error {
	count := 0
	memStorage := storage.NewMemStorage()
	for {
		metric, err := collectMetrics()
		if err != nil {
			return err
		}
		memStorage.AddMetric(metric)
		if count >= flagPollInterval {
			count = 0
			server.SendMetrics(memStorage, flagServAddr)
		}
		time.Sleep(time.Duration(flagReportInterval) * time.Second)
		count += flagReportInterval
	}
}
