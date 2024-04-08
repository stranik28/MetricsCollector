package storage

type Metric struct {
	Gauge   map[string]float64
	Counter uint
}

type MemStorage struct {
	Metrics []Metric
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		Metrics: make([]Metric, 0),
	}
}

func (m *MemStorage) AddMetric(metric Metric) {
	m.Metrics = append(m.Metrics, metric)
}

func CreateMetric(gauge map[string]float64, counter uint) Metric {
	return Metric{Gauge: gauge, Counter: counter}
}

func (m *MemStorage) ClearMemStorage() {
	m.Metrics = m.Metrics[:0]
}
