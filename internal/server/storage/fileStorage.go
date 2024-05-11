package storage

import (
	"encoding/json"
	"os"
)

func SaveMetricsToFile(filename string) {
	metrics, err := GetAll()
	if err != nil {
		return
	}
	data, err := json.Marshal(metrics)
	if err != nil {
		return
	}
	err = os.WriteFile(filename, data, 0666)
	if err != nil {
		return
	}
}

func LoadMetricsFromFile(filename string) (map[string]Metric, error) {
	var metrics map[string]Metric
	data, err := os.ReadFile(filename)
	if err != nil {
		return metrics, err
	}
	err = json.Unmarshal(data, &metrics)
	if err != nil {
		if len(data) == 0 {
			return make(map[string]Metric), nil
		}
		return metrics, err
	}
	return metrics, nil
}
