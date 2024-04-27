package main

import (
	"github.com/stranik28/MetricsCollector/internal/agent"
	"github.com/stranik28/MetricsCollector/internal/agent/collector"
	"log"
)

func main() {
	err := agent.ParsFlags()
	if err != nil {
		panic(err)
	}
	err = collector.MetricsCollector(agent.FlagReportInterval, agent.FlagPollInterval, agent.FlagServAddr)
	if err != nil {
		log.Print("Failed to start metrics collector")
	}
}
