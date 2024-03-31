package storage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemStorage_AddMetric(t *testing.T) {
	type fields struct {
		Metrics []Metric
	}
	type args struct {
		metric Metric
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Simple test #1",
			fields: fields{Metrics: make([]Metric, 0)},
			args: args{metric: Metric{
				Gauge:   map[string]float64{"metric1": 6.32},
				Counter: 1,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemStorage{
				Metrics: tt.fields.Metrics,
			}
			assert.Equal(t, 0, len(m.Metrics))
			m.AddMetric(tt.args.metric)
			assert.Equal(t, 1, len(m.Metrics))
		})
	}
}

func TestMemStorage_ClearMemStorage(t *testing.T) {
	type fields struct {
		Metrics []Metric
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name:   "Simple test #1",
			fields: fields{Metrics: []Metric{{Counter: 1, Gauge: map[string]float64{"some": 1.22}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemStorage{
				Metrics: tt.fields.Metrics,
			}
			assert.Equal(t, 1, len(m.Metrics))
			m.ClearMemStorage()
			assert.Equal(t, 0, len(m.Metrics))
		})
	}
}
