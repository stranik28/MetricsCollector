package storage

type MemStorage struct {
	metrics map[string]Metrics
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		metrics: make(map[string]Metrics),
	}
}

func (mem *MemStorage) SetMemStorage(key string, value Metrics) {
	mem.metrics[key] = value
}

func (mem *MemStorage) GetMemStorage(key string) (Metrics, bool) {
	value, err := mem.metrics[key]
	return value, err
}

func (mem *MemStorage) GetAll() map[string]Metrics {
	return mem.metrics
}
