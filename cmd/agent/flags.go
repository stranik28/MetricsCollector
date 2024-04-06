package main

import (
	"flag"
	"os"
	"strconv"
)

var flagServAddr string
var flagReportInterval int
var flagPollInterval int

func parsFlags() {
	flag.StringVar(&flagServAddr, "a", "127.0.0.1:8080", "address and port to run server")
	flag.IntVar(&flagReportInterval, "r", 2, "Frequency of metrics collecting")
	flag.IntVar(&flagPollInterval, "p", 10, "Frequency of metrics pooling")
	flag.Parse()
	if envRunAddr := os.Getenv("ADDRESS"); envRunAddr != "" {
		flagServAddr = envRunAddr
	}
	if envReportInterval := os.Getenv("REPORT_INTERVAL"); envReportInterval != "" {
		flagReportInterval, _ = strconv.Atoi(envReportInterval)
	}
	if envPollInterval := os.Getenv("POLL_INTERVAL"); envPollInterval != "" {
		flagPollInterval, _ = strconv.Atoi(envPollInterval)
	}
}
