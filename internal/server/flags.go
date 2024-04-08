package server

import (
	"flag"
	"os"
)

var FlagRunAddr string

func ParsFlags() {
	flag.StringVar(&FlagRunAddr, "a", "127.0.0.1:8080", "address and port to run server")
	flag.Parse()
	if envRunAddr := os.Getenv("ADDRESS"); envRunAddr != "" {
		FlagRunAddr = envRunAddr
	}
}
