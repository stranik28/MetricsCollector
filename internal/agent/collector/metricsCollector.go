package collector

import (
	"github.com/stranik28/MetricsCollector/internal/agent/server"
	"github.com/stranik28/MetricsCollector/internal/agent/storage"
	"time"
)

func MetricsCollector(flagReportInterval int, flagPollInterval int, flagServAddr string) {
	count := 0
	memStorage := storage.NewMemStorage()
	for {
		time.Sleep(time.Duration(flagReportInterval) * time.Second)
		count += flagReportInterval
		metric, err := collectMetrics()
		if err != nil {
			panic(err)
		}
		memStorage.AddMetric(metric)
		if count >= flagPollInterval {
			count = 0
			server.SendMetrics(memStorage, flagServAddr)
		}
	}
}
