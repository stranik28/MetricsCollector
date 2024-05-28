package collector

import (
	"encoding/json"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/stranik28/MetricsCollector/internal/agent/storage"
	"go.uber.org/zap"
	"math/rand"
	"runtime"
	"slices"
	"sync"
)

func memStatsToJSON(rtm *runtime.MemStats, inInterface *map[string]interface{}) error {
	inspect, err := json.Marshal(rtm)
	if err != nil {
		return err
	}
	err = json.Unmarshal(inspect, &inInterface)
	if err != nil {
		return err
	}
	return nil
}

func additionalCollect(gauge map[string]float64, mutex *sync.Mutex, logger *zap.Logger) {
	v, err := mem.VirtualMemory()
	if err != nil {
		logger.Error("Error getting memory info: %v", zap.Error(err))
	}
	cpuUtilization, err := cpu.Percent(0, true)
	if err != nil {
		logger.Error("Error getting CPU utilization: %v", zap.Error(err))
	}
	mutex.Lock()
	gauge["TotalMemory"] = float64(v.Total)
	gauge["FreeMemory"] = float64(v.Free)
	gauge["CPUutilization1"] = float64(len(cpuUtilization))
	mutex.Unlock()

}

func collectMetrics(metricsChanel chan storage.Metric, logger *zap.Logger) {
	gauge := make(map[string]float64)
	mutex := sync.Mutex{}
	var counter int64
	go additionalCollect(gauge, &mutex, logger)

	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)
	var inInterface map[string]interface{}

	err := memStatsToJSON(&rtm, &inInterface)

	if err != nil {
		logger.Error("Failed to covertError to Json", zap.Error(err))
	}
	for field, val := range inInterface {
		if slices.Contains(storage.GaugeMetrics, field) {
			if floatValue, ok := val.(float64); ok {
				mutex.Lock()
				gauge[field] = floatValue
				mutex.Unlock()
			}
		}
	}

	gauge["RandomValue"] = rand.Float64()
	counter += 1
	metric := storage.CreateMetric(gauge, counter)
	metricsChanel <- metric
}
