package server

import (
	"bytes"
	"encoding/json"
	"github.com/stranik28/MetricsCollector/internal/agent/models"
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
	client := &http.Client{}
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest(method, serv.url, bytes.NewBuffer(bodyJSON))
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	code := resp.StatusCode

	return code
}
