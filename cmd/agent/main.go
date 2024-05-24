package main

import (
	"fmt"
	"github.com/stranik28/MetricsCollector/internal/agent"
	"github.com/stranik28/MetricsCollector/internal/agent/collector"
	"github.com/stranik28/MetricsCollector/internal/logger"
	"log"
)

func main() {
	err := agent.ParsFlags()
	fmt.Println("Agent parsing")
	if err != nil {
		panic(err)
	}
	loggerV, err := logger.Init("info", "agent.log")
	if err != nil {
		panic(err)
	}
	err = collector.MetricsCollector(agent.FlagReportInterval, agent.FlagPollInterval, agent.FlagServAddr, loggerV, agent.FlagSecretKey)
	if err != nil {
		log.Print("Failed to start metrics collector")
		panic(err)
	}
}
