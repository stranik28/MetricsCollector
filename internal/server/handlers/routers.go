package handlers

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/stranik28/MetricsCollector/internal/agent/logger"
	"github.com/stranik28/MetricsCollector/internal/server/middleware"
)

func Routers() *gin.Engine {
	logger1 := logger.Init("info")
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(middleware.Logger(logger1))
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(middleware.Gzip())
	r.POST("/update/:metricType/:metricName/:metricValue", UpdateMetricsParam)
	r.POST("/update/", UpdateMetrics)
	r.GET("/", AllRecordsHandler)
	r.POST("/value/", GetPostMetric)
	r.GET("/value/:metricType/:metricName", GetMetric)

	return r
}
