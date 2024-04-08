package collector

import (
	"github.com/stranik28/MetricsCollector/internal/agent/storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_collectMetrics(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Test_collectMetrics",
			wantErr: false,
		},
	}
	metricsList := storage.GaugeMetrics
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := collectMetrics()
			if (err != nil) != tt.wantErr {
				t.Errorf("collectMetrics() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, v := range metricsList {
				val, ok := got.Gauge[v]
				assert.Equal(t, ok, true, "%s Not in map", v)
				assert.NotNil(t, val, "%s value is None", v)
			}
			assert.Equal(t, len(metricsList), len(got.Gauge))
		})
	}
}
