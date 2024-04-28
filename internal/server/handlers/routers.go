package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/stranik28/MetricsCollector/internal/server/middleware"
)

func Routers() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(middleware.Logger())
	r.Use(middleware.GzipMiddleware())
	r.POST("/update/:metricType/:metricName/:metricValue", UpdateMetricsParam)
	r.POST("/update/", UpdateMetrics)
	r.GET("/", AllRecordsHandler)
	r.POST("/value/", GetPostMetric)
	r.GET("/value/:metricType/:metricName", GetMetric)

	return r
}
