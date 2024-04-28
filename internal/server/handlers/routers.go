package handlers

import (
	"compress/gzip"
	zipper "github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/stranik28/MetricsCollector/internal/server/middleware"
	"net/http"
)

func Routers() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(middleware.Logger())
	r.Use(zipper.Gzip(zipper.DefaultCompression))
	r.Use(func(c *gin.Context) {
		c.Next()
		if c.Request.Header.Get("Accept-Encoding") == "gzip" {
			c.Header("Content-Encoding", "gzip")
			compress := gzip.NewWriter(c.Writer)
			defer compress.Close()
			compress.Reset(c.Writer)
		}

	})
	r.Use(func(c *gin.Context) {
		// Проверяем наличие заголовка "Content-Encoding" и его значение
		if c.Request.Header.Get("Content-Encoding") == "gzip" {
			// Создаем reader для разархивации тела запроса
			gz, err := gzip.NewReader(c.Request.Body)
			if err != nil {
				// Обработка ошибки, если разархивация невозможна
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			defer gz.Close()

			// Заменяем тело запроса на разархивированное
			c.Request.Body = gz
		}

		// Продолжаем обработку запроса
		c.Next()
	})
	r.POST("/update/:metricType/:metricName/:metricValue", UpdateMetricsParam)
	r.POST("/update/", UpdateMetrics)
	r.GET("/", AllRecordsHandler)
	r.POST("/value/", GetPostMetric)
	r.GET("/value/:metricType/:metricName", GetMetric)

	return r
}
