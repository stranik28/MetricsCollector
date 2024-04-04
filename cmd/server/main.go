package main

import (
	"github.com/gin-gonic/gin"
	handl "github.com/stranik28/MetricsCollector/cmd/server/handlers"
)

func routers() *gin.Engine {
	r := gin.Default()

	r.GET("/update/:metricType/:metricName/:metricValue", handl.UpdateMetrics)
	r.GET("/", handl.AllRecordsHandler)

	return r
}

func main() {
	r := routers()
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
