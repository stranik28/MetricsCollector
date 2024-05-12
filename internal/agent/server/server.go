package server

import (
	"bytes"
	"encoding/json"
	"github.com/stranik28/MetricsCollector/internal/agent/models"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Server struct {
	url string
}

func NewServer(url string) *Server {
	return &Server{
		url: url,
	}
}

func (serv *Server) SendReq(method string, logger *zap.Logger) int {
	client := &http.Client{}
	req, err := http.NewRequest(method, serv.url, nil)
	if err != nil {
		logger.Error("NewRequest: ", zap.Error(err))
	}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Do: ", zap.Error(err))
	}
	defer resp.Body.Close()

	code := resp.StatusCode

	return code
}

func (serv *Server) SendReqPost(method string, body []models.Metrics, logger *zap.Logger) (int, error) {
	retries := []int{1, 3, 5}
	client := &http.Client{}
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		logger.Warn("Can't Marshal Body", zap.Any("body", body))
		return 0, err
	}
	bodyJSONCompressed, err := Compress(bodyJSON, logger)
	if err != nil {
		logger.Error("Error gzip", zap.Any("Error", err.Error()))
		return 0, err
	}

	logger.Info("Sending request", zap.Any("body", string(bodyJSONCompressed)))

	var code int

	for _, i := range retries {
		req, err := http.NewRequest(method, serv.url, bytes.NewBuffer(bodyJSONCompressed))
		if err != nil {
			logger.Fatal("Error" + err.Error())
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Content-Encoding", "gzip")
		resp, err := client.Do(req)
		if err != nil {
			logger.Error("Cant's " + err.Error())
			continue
		}
		defer resp.Body.Close()

		code = resp.StatusCode

		if code >= 200 && code < 300 {
			break
		}
		time.Sleep(time.Duration(i) * time.Second)
	}

	return code, nil
}
