package server

import (
	"fmt"
	"github.com/stranik28/MetricsCollector/internal/server/handlers"
)

func Run() {
	parsFlags()
	r := handlers.Routers()
	fmt.Println("Running server on", flagRunAddr)
	err := r.Run(flagRunAddr)
	if err != nil {
		panic(err)
	}
}
