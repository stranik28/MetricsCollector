package main

import (
	"github.com/stranik28/MetricsCollector/internal/server"
	"github.com/stranik28/MetricsCollector/internal/server/handlers"
	"github.com/stranik28/MetricsCollector/internal/server/storage"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := server.ParsFlags()
	if err != nil {
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
