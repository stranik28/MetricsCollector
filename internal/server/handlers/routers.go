package handlers

import "github.com/gin-gonic/gin"

func Routers() *gin.Engine {
	r := gin.Default()

	r.POST("/update/:metricType/:metricName/:metricValue", UpdateMetrics)
	r.GET("/", AllRecordsHandler)
	r.GET("/value/:metricType/:metricName", GetMetric)

	return r
}
