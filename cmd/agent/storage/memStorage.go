package storage

type MemStorage struct {
	Metrics []Metric
}

func (m *MemStorage) AddMetric(metric Metric) {
	m.Metrics = append(m.Metrics, metric)
}

func (m *MemStorage) ClearMemStorage() {
	m.Metrics = m.Metrics[:0]
}
