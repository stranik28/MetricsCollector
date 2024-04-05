package service

import (
	"github.com/stranik28/MetricsCollector/cmd/server/storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllMetrics(t *testing.T) {
	tests := []struct {
		name    string
		want    map[string]storage.Metric
		wantErr bool
	}{
		{
			name:    "Empty test",
			want:    make(map[string]storage.Metric),
			wantErr: false,
		},
		{
			name: "Simple test #1",
			want: map[string]storage.Metric{"Metric1": {Gauge: 1.89, Counter: 1},
				"metric2": {893482.213914, 9}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.want {
				storage.SetMemStorage(k, v)
			}
			got, err := GetAllMetrics()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllMetrics() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.ObjectsAreEqual(tt.want, got) {
				t.Errorf("GetAllMetrics() got = %v, want %v", got, tt.want)
			}
		})
	}
}
