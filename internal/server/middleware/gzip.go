package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/stranik28/MetricsCollector/internal/server/compress"
	"net/http"
	"strings"
)

func GzipMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// по умолчанию устанавливаем оригинальный http.ResponseWriter как тот,
		// который будем передавать следующей функции
		ow := c.Writer

		// проверяем, что клиент умеет получать от сервера сжатые данные в формате gzip
		acceptEncoding := c.Request.Header.Get("Accept-Encoding")
		supportsGzip := strings.Contains(acceptEncoding, "gzip")
		if supportsGzip {
			// оборачиваем оригинальный http.ResponseWriter новым с поддержкой сжатия
			cw := compress.NewCompressWriter(c.Writer)
			// меняем оригинальный http.ResponseWriter на новый
			ow = cw
			// не забываем отправить клиенту все сжатые данные после завершения middleware
			defer cw.Close()
		}

		// проверяем, что клиент отправил серверу сжатые данные в формате gzip
		contentEncoding := c.Request.Header.Get("Content-Encoding")
		sendsGzip := strings.Contains(contentEncoding, "gzip")
		if sendsGzip {
			// оборачиваем тело запроса в io.Reader с поддержкой декомпрессии
			cr, err := compress.NewCompressReader(c.Request.Body)
			if err != nil {
				c.Writer.WriteHeader(http.StatusInternalServerError)
				return
			}
			// меняем тело запроса на новое
			c.Request.Body = cr
			defer cr.Close()
		}

		// передаём управление хендлеру
		c.Writer = ow
		c.Next()
	}
}
