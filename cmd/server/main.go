package main

import (
	"fmt"
	"github.com/stranik28/MetricsCollector/internal/server"
	"github.com/stranik28/MetricsCollector/internal/server/handlers"
)

func main() {
	server.ParsFlags()
	r := handlers.Routers()
	fmt.Println("Running server on", server.FlagRunAddr)
	err := r.Run(server.FlagRunAddr)
	if err != nil {
		panic(err)
	}
}
