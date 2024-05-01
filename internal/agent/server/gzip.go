package server

import (
	"bytes"
	"compress/gzip"
	"github.com/stranik28/MetricsCollector/internal/agent/logger"
	"go.uber.org/zap"
)

func Compress(data []byte) ([]byte, error) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(data); err != nil {
		logger.Log.Error("Compress error", zap.Error(err))
		return nil, err
	}
	if err := gz.Close(); err != nil {
		logger.Log.Info("Compress error", zap.Error(err))
		return nil, err
	}
	return b.Bytes(), nil
}
