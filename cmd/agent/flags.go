package main

import (
	"flag"
)

var flagServAddr string
var flagReportInterval int
var flagPollInterval int

func parsFlags() {
	flag.StringVar(&flagServAddr, "a", "127.0.0.1:8080", "address and port to run server")
	flag.IntVar(&flagReportInterval, "r", 2, "Frequency of metrics collecting")
	flag.IntVar(&flagPollInterval, "p", 10, "Frequency of metrics pooling")
	flag.Parse()
}
