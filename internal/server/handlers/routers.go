package handlers

import (
	"bytes"
	"compress/gzip"
	gingzip "github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/stranik28/MetricsCollector/internal/server/logger"
	"github.com/stranik28/MetricsCollector/internal/server/middleware"
	"go.uber.org/zap"
	"net/http"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func Routers() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(middleware.Logger())
	r.Use(gingzip.Gzip(gzip.DefaultCompression))
	r.Use(func(c *gin.Context) {
		c.Next()
		if c.Request.Header.Get("Accept-Encoding") == "gzip" {
			//	logger.Log.Warn("Accepting Gzip")
			//	c.Header("Content-Encoding", "gzip")
			//	logger.Log.Warn("Accepting Gzip")
			//	compress := gzip.NewWriter(c.Writer)
			//	logger.Log.Warn("Compressing Gzip")
			//	defer compress.Close()
			//	compress.Reset(c.Writer)
			//	logger.Log.Warn("Compressing Reset")
			w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
			c.Writer = w
			c.Next()
			logger.Log.Info("Response body: ", zap.Any("Data", w.body.Bytes))
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
