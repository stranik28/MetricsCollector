package main

import (
	"github.com/stranik28/MetricsCollector/internal/server"
	"github.com/stranik28/MetricsCollector/internal/server/handlers"
	"github.com/stranik28/MetricsCollector/internal/server/logger"
	"go.uber.org/zap"
)

func main() {
	server.ParsFlags()
	err := logger.Init("info")
	if err != nil {
		panic(err)
	}
	r := handlers.Routers()
	logger.Log.Info("Running server", zap.String("address", server.FlagRunAddr))
	err = r.Run(server.FlagRunAddr)
	if err != nil {
		panic(err)
	}
}
