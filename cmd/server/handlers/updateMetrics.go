package handlers

import (
	stor "github.com/stranik28/MetricsCollector/cmd/server/storage"
	"net/http"
	"strconv"
	"strings"
)

var storage = stor.NewMemStorage()

func UpdateMetrics(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(res, "METHOD NOT ALLOWED "+req.Method, http.StatusMethodNotAllowed)
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
		val, ok := storage.GetMemStorage(params[2])
		if !ok {
			val = stor.Metrics{
				Gauge:   0,
				Counter: value,
			}
		} else {
			val.Counter += value
		}
		storage.SetMemStorage(params[2], val)
	case "gauge":
		value, err := strconv.ParseFloat(strings.TrimSpace(params[3]), 64)
		if err != nil {
			http.Error(res, "Значение должно быть в формате float64", http.StatusBadRequest)
			return
		}
		val, ok := storage.GetMemStorage(params[2])
		if !ok {
			val = stor.Metrics{
				Gauge:   value,
				Counter: 0,
			}
		} else {
			val.Gauge = value
		}
		storage.SetMemStorage(params[2], val)
	default:
		http.Error(res, "Недопустимый тип метрики", http.StatusBadRequest)
		return
	}

	res.WriteHeader(http.StatusOK)
}
