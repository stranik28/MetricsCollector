package main

import (
	"github.com/stranik28/MetricsCollector/internal/agent"
	"github.com/stranik28/MetricsCollector/internal/agent/collector"
	"github.com/stranik28/MetricsCollector/internal/agent/logger"
	"log"
)

func main() {
	err := agent.ParsFlags()
	if err != nil {
		panic(err)
	}
	loggerV, err := logger.Init("info")
	if err != nil {
		panic(err)
	}
	err = collector.MetricsCollector(agent.FlagReportInterval, agent.FlagPollInterval, agent.FlagServAddr, loggerV)
	if err != nil {
		log.Print("Failed to start metrics collector")
		panic(err)
	}
}
