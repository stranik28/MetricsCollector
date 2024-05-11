package server

import (
	"flag"
	"os"
	"strconv"
)

var FlagRunAddr string
var StoreInterval int
var FileStoragePath string
var Restore bool
var DBURL string

func ParsFlags() error {
	flag.StringVar(&FlagRunAddr, "a", "127.0.0.1:8080", "address and port to run server")
	flag.IntVar(&StoreInterval, "i", 300, "interval before save settings")
	flag.StringVar(&FileStoragePath, "f", "/tmp/metrics-db.json", "location of metrics db file")
	flag.BoolVar(&Restore, "r", true, "restore metrics db")
	flag.StringVar(&DBURL, "d", "", "database url")
	flag.Parse()
	var err error
	if envRunAddr := os.Getenv("ADDRESS"); envRunAddr != "" {
		FlagRunAddr = envRunAddr
	}
	if envStoreInterval := os.Getenv("STORE_INTERVAL"); envStoreInterval != "" {
		StoreInterval, err = strconv.Atoi(envStoreInterval)
		if err != nil {
			return err
		}
	}
	if envFileStoragePath := os.Getenv("FILE_STORAGE_PATH"); envFileStoragePath != "" {
		FileStoragePath = envFileStoragePath
	}
	if envRestore := os.Getenv("RESTORE"); envRestore != "" {
		Restore, err = strconv.ParseBool(envRestore)
		if err != nil {
			return err
		}
	}
	if envDBURL := os.Getenv("DB_URL"); envDBURL != "" {
		DBURL = envDBURL
	}
	return nil
}
