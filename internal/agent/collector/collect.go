package collector

import (
	"encoding/json"
	"fmt"
	"github.com/stranik28/MetricsCollector/internal/agent/storage"
	"math/rand"
	"runtime"
	"slices"
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

func collectMetrics() (storage.Metric, error) {
	gauge := make(map[string]float64)

	var counter int64

	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)
	var inInterface map[string]interface{}

	err := memStatsToJSON(&rtm, &inInterface)

	if err != nil {
		return storage.Metric{}, err
	}

	for field, val := range inInterface {
		if slices.Contains(storage.GaugeMetrics, field) {
			if floatValue, ok := val.(float64); ok {
				gauge[field] = floatValue
			} else {
				fmt.Printf("Failed to convert %v to float64\n", field)
			}
		}
	}
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	gauge["RandomValue"] = rand.Float64()
	counter += 1
	metric := storage.CreateMetric(gauge, counter)

	return metric, nil
}
