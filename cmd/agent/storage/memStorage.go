package storage

type MemStorage struct {
	Metrics []Metrics
}

func (m *MemStorage) AddMetric(metric Metrics) {
	m.Metrics = append(m.Metrics, metric)
}

func (m *MemStorage) ClearMemStorage() {
	m.Metrics = m.Metrics[:0]
}
