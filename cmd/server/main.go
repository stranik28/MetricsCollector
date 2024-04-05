package main

import "github.com/stranik28/MetricsCollector/cmd/server/handlers"

func main() {
	r := handlers.Routers()
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
