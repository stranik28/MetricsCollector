package middleware

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

// Signature создает HMAC-SHA256 подпись
func Signature(body []byte, key []byte) []byte {
	hash := hmac.New(sha256.New, key)
	hash.Write(body)
	return hash.Sum(nil)
}

func SignatureMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		signature := c.GetHeader("Hash")

		if secretKey != "" && signature != "" && signature != "none" {
			signatureHex, err := hex.DecodeString(signature)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get Signature"})
				c.Abort()
				return
			}
			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err != nil {

				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read body"})
				c.Abort()
				return
			}

			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			expectedSignature := Signature(bodyBytes, []byte(secretKey))
			if !hmac.Equal(signatureHex, expectedSignature) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid signature"})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}

func ResponseSignatureMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Создаем буфер для записи ответа
		bodyWriter := &responseBodyWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter

		c.Next()

		// После вызова следующего обработчика и записи ответа
		responseBody := bodyWriter.body.Bytes()
		signature := Signature(responseBody, []byte(secretKey))

		// Добавляем подпись в заголовок ответа
		c.Writer.Header().Set("HashSHA256", hex.EncodeToString(signature))
	}
}

// responseBodyWriter - структура для перехвата тела ответа
type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *responseBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
