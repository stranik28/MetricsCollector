package main

import (
	"fmt"
	"github.com/stranik28/MetricsCollector/cmd/server/handlers"
)

func main() {
	parsFlags()
	r := handlers.Routers()
	fmt.Println("Running server on", flagRunAddr)
	err := r.Run(flagRunAddr)
	if err != nil {
		panic(err)
	}
}
