package main

import (
	"github.com/stranik28/MetricsCollector/internal/server"
	"github.com/stranik28/MetricsCollector/internal/server/handlers"
	"github.com/stranik28/MetricsCollector/internal/server/logger"
	"github.com/stranik28/MetricsCollector/internal/server/storage"
	"go.uber.org/zap"
)

func main() {
	err := logger.Init("info")
	if err != nil {
		panic(err)
	}
	err = server.ParsFlags()
	if err != nil {
		logger.Log.Info("Failed to parse flags", zap.Error(err))
		panic(err)
	}
	go storage.InitFileSave(server.FileStoragePath, server.Restore, server.StoreInterval)
	r := handlers.Routers()
	logger.Log.Info("Running server", zap.String("address", server.FlagRunAddr))
	err = r.Run(server.FlagRunAddr)
	if err != nil {
		panic(err)
	}
}
