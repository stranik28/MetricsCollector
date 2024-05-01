package main

import (
	"github.com/stranik28/MetricsCollector/internal/server"
	"github.com/stranik28/MetricsCollector/internal/server/handlers"
	"github.com/stranik28/MetricsCollector/internal/server/middleware"
	"github.com/stranik28/MetricsCollector/internal/server/storage"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := middleware.Init("info")
	if err != nil {
		panic(err)
	}
	err = server.ParsFlags()
	if err != nil {
		//logger.Log.Info("Failed to parse flags", zap.Error(err))
		panic(err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go storage.InitFileSave(server.FileStoragePath, server.Restore, server.StoreInterval, done)
	r := handlers.Routers()
	err = r.Run(server.FlagRunAddr)
	if err != nil {
		panic(err)
	}
}
