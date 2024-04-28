package server

import (
	"flag"
	"github.com/stranik28/MetricsCollector/internal/server/logger"
	"os"
	"strconv"
	"time"
)

var FlagRunAddr string
var StoreInterval time.Duration
var FileStoragePath string
var Restore bool

func ParsFlags() error {
	flag.StringVar(&FlagRunAddr, "a", "127.0.0.1:8080", "address and port to run server")
	flag.DurationVar(&StoreInterval, "i", 300, "interval before save settings")
	flag.StringVar(&FileStoragePath, "f", "/tmp/metrics-db.json", "location of metrics db file")
	flag.BoolVar(&Restore, "r", true, "restore metrics db")
	flag.Parse()
	var err error
	if envRunAddr := os.Getenv("ADDRESS"); envRunAddr != "" {
		FlagRunAddr = envRunAddr
	}
	if envStoreInterval := os.Getenv("STORE_INTERVAL"); envStoreInterval != "" {
		StoreInterval, err = time.ParseDuration(envStoreInterval)
		if err != nil {
			logger.Log.Debug("Error parsing STORE_INTERVAL env var")
			return err
		}
	}
	if envFileStoragePath := os.Getenv("FILE_STORAGE_PATH"); envFileStoragePath != "" {
		FileStoragePath = envFileStoragePath
	}
	if envRestore := os.Getenv("RESTORE"); envRestore != "" {
		Restore, err = strconv.ParseBool(envRestore)
		if err != nil {
			logger.Log.Debug("Error parsing RESTORE env var")
			return err
		}
	}
	return nil
}
