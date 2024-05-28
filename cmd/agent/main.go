package main

import (
	"github.com/stranik28/MetricsCollector/internal/agent"
	"github.com/stranik28/MetricsCollector/internal/agent/collector"
	"github.com/stranik28/MetricsCollector/internal/logger"
)

func main() {
	err := agent.ParsFlags()
	if err != nil {
		panic(err)
	}
	loggerV, err := logger.Init("info", "agent.log")
	if err != nil {
		panic(err)
	}
	collector.MetricsCollector(agent.FlagReportInterval, agent.FlagPollInterval, agent.FlagServAddr, loggerV, agent.FlagSecretKey, agent.RateLimit)
}
