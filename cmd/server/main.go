package main

import (
	"net/http"
	"strconv"
	"strings"
)

type Metrics struct {
	gauge   float64
	counter int64
}

type MemStorage struct {
	metrics map[string]Metrics
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		metrics: make(map[string]Metrics),
	}
}

func (mem *MemStorage) setMemStorage(key string, value Metrics) {
	mem.metrics[key] = value
}

func (mem *MemStorage) getMemStorage(key string) (Metrics, bool) {
	value, err := mem.metrics[key]
	return value, err
}

func updateMetrics(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(res, "METHOD NOT ALLOWED", http.StatusMethodNotAllowed)
		return
	}
	url := req.URL.Path
	params := strings.Split(url, "/")
	params = params[1:]
	if len(params) != 4 {
		http.Error(res, "Недопустимое количество параметров", http.StatusNotFound)
		return
	}
	switch params[1] {
	case "counter":
		value, err := strconv.ParseInt(strings.TrimSpace(params[3]), 10, 64)
		if err != nil {
			http.Error(res, "Значение должно быть в формате int64", http.StatusBadRequest)
			return
		}
		val, ok := storage.getMemStorage(params[2])
		if !ok {
			val = Metrics{
				gauge:   0,
				counter: value,
			}
		} else {
			val.counter += value
		}
		storage.setMemStorage(params[2], val)
	case "gauge":
		value, err := strconv.ParseFloat(strings.TrimSpace(params[3]), 64)
		if err != nil {
			http.Error(res, "Значение должно быть в формате float64", http.StatusBadRequest)
			return
		}
		val, ok := storage.getMemStorage(params[2])
		if !ok {
			val = Metrics{
				gauge:   value,
				counter: 0,
			}
		} else {
			val.gauge = value
		}
		storage.setMemStorage(params[2], val)
	default:
		http.Error(res, "Недопустимый тип метрики", http.StatusBadRequest)
		return
	}

	res.WriteHeader(http.StatusOK)
}

var storage = NewMemStorage()

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/update/", updateMetrics)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
