package storage

type Metric struct {
	Gauge   map[string]float64
	Counter uint
}
