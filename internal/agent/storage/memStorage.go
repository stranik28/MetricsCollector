package storage

type Metric struct {
	Gauge   map[string]float64
	Counter int64
}

func CreateMetric(gauge map[string]float64, counter int64) Metric {
	return Metric{Gauge: gauge, Counter: counter}
}
