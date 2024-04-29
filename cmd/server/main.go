package main

import (
	"github.com/stranik28/MetricsCollector/internal/server"
	"github.com/stranik28/MetricsCollector/internal/server/handlers"
	"github.com/stranik28/MetricsCollector/internal/server/logger"
	"github.com/stranik28/MetricsCollector/internal/server/storage"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
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

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go storage.InitFileSave(server.FileStoragePath, server.Restore, server.StoreInterval, done)
	r := handlers.Routers()
	logger.Log.Info("Running server", zap.String("address", server.FlagRunAddr))
	err = r.Run(server.FlagRunAddr)
	if err != nil {
		panic(err)
	}
}
