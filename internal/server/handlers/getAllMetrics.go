package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/stranik28/MetricsCollector/internal/server/service"
	"log"
	"net/http"
)

func AllRecordsHandler(c *gin.Context) {
	metrics, err := service.GetAllMetrics()
	if err != nil {
		log.Printf("Error getting metrics: %v", err)
	}
	c.JSON(http.StatusOK, metrics)
}
