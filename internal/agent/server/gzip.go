package server

import (
	"bytes"
	"compress/gzip"
	"go.uber.org/zap"
)

func Compress(data []byte, logger *zap.Logger) ([]byte, error) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(data); err != nil {
		logger.Error("Compress error", zap.Error(err))
		return nil, err
	}
	if err := gz.Close(); err != nil {
		logger.Info("Compress error", zap.Error(err))
		return nil, err
	}
	return b.Bytes(), nil
}
