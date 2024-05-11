package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/stranik28/MetricsCollector/internal/server"
	"github.com/stranik28/MetricsCollector/internal/server/storage"
	"net/http"
)

func CheckDBConnection(c *gin.Context) {
	err := storage.Connect(server.DBURL)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, nil)
}
