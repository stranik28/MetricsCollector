package storage

import (
	"encoding/json"
	"errors"
	"os"
	"time"
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
	var pathError *os.PathError
	if errors.As(err, &pathError) {
		retries := []int{1, 3, 5}
		for _, v := range retries {
			if err := os.WriteFile(filename, data, 0666); err != nil {
				if errors.As(err, &pathError) {
					time.Sleep(time.Duration(v) * time.Millisecond)
				} else {
					panic(err)
				}
			} else {
				return
			}
		}
	} else if err != nil {
		panic(err)
	}
}

func LoadMetricsFromFile(filename string) (map[string]Metric, error) {
	var metrics map[string]Metric
	data, err := os.ReadFile(filename)
	var pathError *os.PathError
	if errors.As(err, &pathError) {
		retries := []int{1, 3, 5}
		for _, v := range retries {
			if errors.As(err, &pathError) {
				time.Sleep(time.Duration(v) * time.Second)
				data, err = os.ReadFile(filename)
			} else if err != nil {
				return nil, err
			} else {
				break
			}
		}
	} else if err != nil {
		return nil, err
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
