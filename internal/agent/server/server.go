package server

import (
	"bytes"
	"encoding/json"
	"github.com/stranik28/MetricsCollector/internal/agent/logger"
	"github.com/stranik28/MetricsCollector/internal/agent/models"
	"go.uber.org/zap"
	"log"
	"net/http"
)

type Server struct {
	url string
}

func NewServer(url string) *Server {
	return &Server{
		url: url,
	}
}

func (serv *Server) SendReq(method string) int {
	client := &http.Client{}
	req, err := http.NewRequest(method, serv.url, nil)
	if err != nil {
		log.Print("NewRequest: ", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Print("Do: ", err)
	}
	defer resp.Body.Close()

	code := resp.StatusCode

	return code
}

func (serv *Server) SendReqPost(method string, body models.Metrics) int {
	maxRetries := 10
	client := &http.Client{}
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		logger.Log.Warn("Can't Marshal Body", zap.Any("body", body))
	}
	bodyJSONCompressed, err := Compress(bodyJSON)
	if err != nil {
		logger.Log.Error("Error gzip", zap.Any("Error", err.Error()))
	}

	logger.Log.Info("Sending request", zap.Any("body", string(bodyJSONCompressed)))

	var code int

	for i := 0; i < maxRetries; i++ {
		req, err := http.NewRequest(method, serv.url, bytes.NewBuffer(bodyJSONCompressed))
		if err != nil {
			logger.Log.Fatal("Error" + err.Error())
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Content-Encoding", "gzip")
		resp, err := client.Do(req)
		if err != nil {
			logger.Log.Error("Cant's " + err.Error())
			continue
		}
		defer resp.Body.Close()

		code = resp.StatusCode

		if code >= 200 && code < 300 {
			break
		}
	}

	return code
}
