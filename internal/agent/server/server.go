package server

import "net/http"

type Server struct {
	url    string
	method string
}

func NewRequest(url string, method string) *Server {
	return &Server{
		url:    url,
		method: method,
	}
}

func (serv *Server) SendReq() int {
	client := &http.Client{}
	req, err := http.NewRequest(serv.method, serv.url, nil)
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
