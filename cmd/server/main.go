package main

import (
	handl "github.com/stranik28/MetricsCollector/cmd/server/handlers"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/update/", handl.UpdateMetrics)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
