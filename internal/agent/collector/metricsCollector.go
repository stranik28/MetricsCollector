package collector

import (
	"github.com/stranik28/MetricsCollector/internal/agent/server"
	"github.com/stranik28/MetricsCollector/internal/agent/storage"
	"go.uber.org/zap"
	"sync"
	"time"
)

func MetricsCollector(flagReportInterval int, flagPollInterval int, flagServAddr string, logger *zap.Logger, secretKey string, rateLimit int) {
	metricsChanel := make(chan storage.Metric)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		defer close(metricsChanel)
		for {
			logger.Info("Collecting metrics...")
			go collectMetrics(metricsChanel, logger)
			time.Sleep(time.Duration(flagReportInterval) * time.Second)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			logger.Info("Polling metrics...")
			go server.SendMetrics(metricsChanel, flagServAddr, logger, secretKey)
			time.Sleep(time.Duration(flagPollInterval) * time.Second)
		}
	}()
	wg.Wait()
}
