package handlers

import (
	"github.com/gin-gonic/gin"
	stor "github.com/stranik28/MetricsCollector/cmd/server/storage"
	"net/http"
	"strconv"
	"strings"
)

var storage = stor.NewMemStorage()

//func updateMetrics(metricType string, metricValue string, metricName string){
//	switch metricType {
//	case "counter":
//		value, err := strconv.ParseInt(strings.TrimSpace(metricValue), 10, 64)
//		if err != nil {
//			c.JSON(http.StatusBadRequest, "Значение должно быть в формате int64")
//			return
//		}
//		val, ok := storage.GetMemStorage(metricName)
//		if !ok {
//			val = stor.Metrics{
//				Gauge:   0,
//				Counter: value,
//			}
//		} else {
//			val.Counter += value
//		}
//		storage.SetMemStorage(metricName, val)
//	case "gauge":
//		value, err := strconv.ParseFloat(strings.TrimSpace(metricValue), 64)
//		if err != nil {
//			c.JSON(http.StatusBadRequest, "Значение должно быть в формате float64")
//			return
//		}
//		val, ok := storage.GetMemStorage(metricName)
//		if !ok {
//			val = stor.Metrics{
//				Gauge:   value,
//				Counter: 0,
//			}
//		} else {
//			val.Gauge = value
//		}
//		storage.SetMemStorage(metricName, val)
//	default:
//		c.JSON(http.StatusBadRequest, "Недопустимый тип метрики")
//		return
//	}
//
//	c.JSON(http.StatusOK, nil)
//}

func UpdateMetrics(c *gin.Context) {
	metricType := c.Param("metricType")
	metricName := c.Param("metricName")
	metricValue := c.Param("metricValue")
	switch metricType {
	case "counter":
		value, err := strconv.ParseInt(strings.TrimSpace(metricValue), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Значение должно быть в формате int64")
			return
		}
		val, ok := storage.GetMemStorage(metricName)
		if !ok {
			val = stor.Metrics{
				Gauge:   0,
				Counter: value,
			}
		} else {
			val.Counter += value
		}
		storage.SetMemStorage(metricName, val)
	case "gauge":
		value, err := strconv.ParseFloat(strings.TrimSpace(metricValue), 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Значение должно быть в формате float64")
			return
		}
		val, ok := storage.GetMemStorage(metricName)
		if !ok {
			val = stor.Metrics{
				Gauge:   value,
				Counter: 0,
			}
		} else {
			val.Gauge = value
		}
		storage.SetMemStorage(metricName, val)
	default:
		c.JSON(http.StatusBadRequest, "Недопустимый тип метрики")
		return
	}

	c.JSON(http.StatusOK, nil)
}

func AllRecordsHandler(c *gin.Context) {
	metrics := storage.GetAll()
	c.JSON(http.StatusOK, metrics)
}

func GetMetric(c *gin.Context) {
	metricName := c.Param("metricName")
	metricType := c.Param("metricType")
	metric, ok := storage.GetMemStorage(metricName)
	if !ok {
		c.JSON(http.StatusNotFound, "Недопустимое имя")
	}
	if metricType == "guage" {
		c.JSON(http.StatusOK, metric.Gauge)
	} else {
		c.JSON(http.StatusOK, metric.Counter)
	}
}
