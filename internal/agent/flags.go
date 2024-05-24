package agent

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var FlagServAddr string
var FlagReportInterval int
var FlagPollInterval int
var FlagSecretKey string

func ParsFlags() error {

	flag.StringVar(&FlagServAddr, "a", "127.0.0.1:8080", "address and port to run server")
	flag.IntVar(&FlagReportInterval, "r", 2, "Frequency of metrics collecting")
	flag.IntVar(&FlagPollInterval, "p", 10, "Frequency of metrics pooling")
	flag.StringVar(&FlagSecretKey, "k", "", "Secret key")

	flag.Parse()

	var err error

	if envRunAddr := os.Getenv("ADDRESS"); envRunAddr != "" {
		FlagServAddr = envRunAddr
	}

	if envReportInterval := os.Getenv("REPORT_INTERVAL"); envReportInterval != "" {
		FlagReportInterval, err = strconv.Atoi(envReportInterval)
		if err != nil {
			return err
		}
	}

	if envPollInterval := os.Getenv("POLL_INTERVAL"); envPollInterval != "" {
		FlagPollInterval, err = strconv.Atoi(envPollInterval)
		if err != nil {
			return err
		}
	}

	if envSecretKey := os.Getenv("KEY"); envSecretKey != "" {
		FlagSecretKey = envSecretKey
	}
	fmt.Println("Secret key:", FlagSecretKey)
	return nil
}
