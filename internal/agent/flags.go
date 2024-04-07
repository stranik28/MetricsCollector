package agent

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

	var err error

	if envRunAddr := os.Getenv("ADDRESS"); envRunAddr != "" {
		flagServAddr = envRunAddr
	}

	if envReportInterval := os.Getenv("REPORT_INTERVAL"); envReportInterval != "" {
		flagReportInterval, err = strconv.Atoi(envReportInterval)
		if err != nil {
			panic(err)
		}
	}

	if envPollInterval := os.Getenv("POLL_INTERVAL"); envPollInterval != "" {
		flagPollInterval, err = strconv.Atoi(envPollInterval)
		if err != nil {
			panic(err)
		}
	}
}
