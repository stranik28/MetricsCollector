package service

import (
	"context"
	"github.com/stranik28/MetricsCollector/internal/logger"
	"github.com/stranik28/MetricsCollector/internal/server/models"
	"github.com/stranik28/MetricsCollector/internal/server/storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMetricByName(t *testing.T) {
	testLogger, err := logger.Init("info", "test.log")
	assert.Nil(t, err)
	tests := []struct {
		name    string
		args    models.Metrics
		want    interface{}
		wantErr error
	}{
		{
			name: "Positive gauge",
			args: models.Metrics{
				ID:    "Gauge metric",
				MType: "gauge",
			},
			want:    6.66,
			wantErr: nil,
		},
		{
			name: "Positive counter",
			args: models.Metrics{
				ID:    "Counter metric",
				MType: "counter",
			},
			want:    1,
			wantErr: nil,
		},
		{
			name: "Negative test (Fake name)",
			args: models.Metrics{
				ID:    "non exist",
				MType: "counter",
			},
			want:    1,
			wantErr: storage.ErrorMetricsNotFound,
		},
		{
			name: "Negative test (Fake metric type)",
			args: models.Metrics{
				ID:    "Gauge metric",
				MType: "Cringe",
			},
			want:    1,
			wantErr: storage.ErrorMetricsNotFound,
		},
	}
	for k, v := range map[string]storage.Metric{"Gauge metric": {Gauge: 6.66, Counter: 1},
		"Counter metric": {Gauge: 893482.213914, Counter: 1}} {
		storage.SetMemStorage(k, v)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := GetMetricByName(context.Background(), tt.args, testLogger)
			if tt.wantErr == nil {
				if got.MType == "counter" {
					assert.EqualValuesf(t, tt.want, *got.Delta, "GetMetricByName(%v, %v)", tt.args.ID, tt.args.MType)
				} else {
					assert.EqualValuesf(t, tt.want, *got.Value, "GetMetricByName(%v, %v)", tt.args.ID, tt.args.MType)
				}
			} else {
				assert.EqualError(t, err, tt.wantErr.Error())
			}
		})
	}
}

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
				"metric2": {Gauge: 893482.213914, Counter: 9}},
			wantErr: false,
		},
	}
	storage.ClearStorage()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.want {
				storage.SetMemStorage(k, v)
			}
			got, err := GetAllMetrics(context.Background())
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

func TestUpdateMetrics(t *testing.T) {
	val1 := 6.66
	val2 := 7.01
	var delta1 int64 = 6
	var delta2 int64 = 98
	var delta3 int64 = 123
	tests := []struct {
		name    string
		args    models.Metrics
		wantErr error
	}{
		{
			name: "Positive Insert new gauge",
			args: models.Metrics{
				MType: "gauge",
				Value: &val1,
				ID:    "New gauge",
			},
			wantErr: nil,
		},
		{
			name: "Positive Update gauge",
			args: models.Metrics{
				MType: "gauge",
				Value: &val2,
				ID:    "New gauge",
			},
			wantErr: nil,
		},
		{
			name: "Positive Insert new counter",
			args: models.Metrics{
				MType: "counter",
				Delta: &delta1,
				ID:    "New counter",
			},
			wantErr: nil,
		},
		{
			name: "Positive Update counter",
			args: models.Metrics{
				MType: "counter",
				Delta: &delta2,
				ID:    "New counter",
			},
			wantErr: nil,
		},
		{
			name: "Negative test #1",
			args: models.Metrics{
				MType: "idk",
				Delta: &delta3,
				ID:    "New idk",
			},
			wantErr: storage.ErrorIncorrectTypeMetrics,
		},
	}
	storage.ClearStorage()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			batch := make([]models.Metrics, 1)
			batch[0] = tt.args
			_, err := UpdateMetrics(context.Background(), batch)
			if err != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
			}
		})
	}
}
