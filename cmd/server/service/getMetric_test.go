package service

import (
	"github.com/stranik28/MetricsCollector/cmd/server/storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMetricByName(t *testing.T) {
	type args struct {
		metricName string
		metricType string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr error
	}{
		{
			name: "Positive gauge",
			args: args{
				metricName: "Gauge metric",
				metricType: "gauge",
			},
			want:    6.66,
			wantErr: nil,
		},
		{
			name: "Positive counter",
			args: args{
				metricName: "Counter metric",
				metricType: "counter",
			},
			want:    1,
			wantErr: nil,
		},
		{
			name: "Negative test (Fake name)",
			args: args{
				metricName: "non exist",
				metricType: "counter",
			},
			want:    1,
			wantErr: storage.ErrorMetricsNotFound,
		},
		{
			name: "Negative test (Fake metric type)",
			args: args{
				metricName: "Gauge metric",
				metricType: "Cringe",
			},
			want:    1,
			wantErr: storage.ErrorMetricsNotFound,
		},
	}
	for k, v := range map[string]storage.Metric{"Gauge metric": {6.66, 1},
		"Counter metric": {893482.213914, 1}} {
		storage.SetMemStorage(k, v)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMetricByName(tt.args.metricName, tt.args.metricType)
			if tt.wantErr == nil {
				assert.EqualValuesf(t, tt.want, got, "GetMetricByName(%v, %v)", tt.args.metricName, tt.args.metricType)
			} else {
				assert.EqualError(t, err, tt.wantErr.Error())
			}
		})
	}
}
