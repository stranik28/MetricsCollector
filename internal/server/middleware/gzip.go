package middleware

import (
	"compress/gzip"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Gzip() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8; text/html; charset=utf-8")
		if c.Request.Header.Get("Content-Encoding") == "gzip" {
			gz, err := gzip.NewReader(c.Request.Body)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			defer gz.Close()
			c.Request.Body = gz
		}
		c.Next()
	}
}
