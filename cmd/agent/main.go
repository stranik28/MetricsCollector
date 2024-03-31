package main

import (
	"encoding/json"
	"fmt"
	"github.com/stranik28/MetricsCollector/cmd/agent/server"
	"github.com/stranik28/MetricsCollector/cmd/agent/storage"
	"math/rand/v2"
	"runtime"
	"slices"
	"time"
)

var GaugeMetrics = []string{
	"Alloc",
	"BuckHashSys",
	"Frees",
	"GCCPUFraction",
	"GCSys",
	"HeapAlloc",
	"HeapIdle",
	"HeapInuse",
	"HeapObjects",
	"HeapReleased",
	"HeapSys",
	"LastGC",
	"Lookups",
	"MCacheInuse",
	"MCacheSys",
	"MSpanInuse",
	"MSpanSys",
	"Mallocs",
	"NextGC",
	"NumForcedGC",
	"NumGC",
	"OtherSys",
	"PauseTotalNs",
	"StackInuse",
	"StackSys",
	"Sys",
	"TotalAlloc",
	"RandomValue",
}

func collectMetrics() (map[string]float64, uint, error) {
	gauge := make(map[string]float64)

	var counter uint

	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)
	var inInterface map[string]interface{}

	inrec, _ := json.Marshal(rtm)
	err := json.Unmarshal(inrec, &inInterface)
	if err != nil {
		return nil, 0, err
	}
	for field, val := range inInterface {
		if slices.Contains(GaugeMetrics, field) {
			if floatValue, ok := val.(float64); ok {
				gauge[field] = floatValue
			} else {
				fmt.Printf("Failed to convert %v to float64\n", field)
			}
		}
	}
	gauge["RandomValue"] = rand.Float64()
	counter += 1
	return gauge, counter, nil
}

func main() {
	count := 0
	memStorage := storage.MemStorage{Metrics: make([]storage.Metrics, 0)}
	for {
		time.Sleep(2 * time.Second)
		count += 1
		gauge, counter, err := collectMetrics()
		metric := storage.Metrics{Gauge: gauge, Counter: counter}
		memStorage.AddMetric(metric)
		if err != nil {
			panic(err)
		}
		if count == 5 {
			count = 0
			for _, stor := range memStorage.Metrics {
				for k, v := range stor.Gauge {
					url := fmt.Sprintf("http://localhost:8080/update/gauge/%s/%f", k, v)
					server.SendReq(url, "POST")
				}
				url := fmt.Sprintf("http://localhost:8080/update/counter/PollCount/%d", stor.Counter)
				server.SendReq(url, "POST")

			}
			memStorage.ClearMemStorage()
		}
	}

}
