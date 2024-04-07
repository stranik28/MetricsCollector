package agent

import "github.com/stranik28/MetricsCollector/internal/agent/collector"

func Run() {
	parsFlags()
	collector.MetricsCollector(flagReportInterval, flagPollInterval, flagServAddr)
}
