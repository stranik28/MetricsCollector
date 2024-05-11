package handlers

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/stranik28/MetricsCollector/internal/logger"
	"github.com/stranik28/MetricsCollector/internal/server/middleware"
)

func Routers() *gin.Engine {
	logger1, err := logger.Init("info", "server.log")
	if err != nil {
		panic(err)
	}
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(middleware.Logger(logger1))
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(middleware.Gzip())
	r.POST("/update/:metricType/:metricName/:metricValue", UpdateMetricsParam)
	r.POST("/update/", UpdateMetrics)
	r.POST("/updates/", UpdateMetrics)
	r.GET("/", AllRecordsHandler)
	r.POST("/value/", GetPostMetric)
	r.GET("/value/:metricType/:metricName", GetMetric)
	r.GET("/ping", CheckDBConnection)

	return r
}
