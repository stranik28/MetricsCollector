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
	for k, v := range map[string]storage.Metric{"Gauge metric": {Gauge: 6.66, Counter: 1},
		"Counter metric": {Gauge: 893482.213914, Counter: 1}} {
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

func TestUpdateMetrics(t *testing.T) {
	type args struct {
		metricType  string
		metricValue string
		metricName  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Positive Insert new gauge",
			args: args{
				metricType:  "gauge",
				metricValue: "6.98584",
				metricName:  "New gauge",
			},
			wantErr: nil,
		},
		{
			name: "Positive Update gauge",
			args: args{
				metricType:  "gauge",
				metricValue: "7.01",
				metricName:  "New gauge",
			},
			wantErr: nil,
		},
		{
			name: "Negative Update gauge",
			args: args{
				metricType:  "gauge",
				metricValue: "seven point five",
				metricName:  "New gauge",
			},
			wantErr: storage.ErrorIncorrectTypeFloat64,
		},
		{
			name: "Positive Insert new counter",
			args: args{
				metricType:  "counter",
				metricValue: "6",
				metricName:  "New counter",
			},
			wantErr: nil,
		},
		{
			name: "Positive Update counter",
			args: args{
				metricType:  "counter",
				metricValue: "98",
				metricName:  "New counter",
			},
			wantErr: nil,
		},
		{
			name: "Negative Update gauge",
			args: args{
				metricType:  "counter",
				metricValue: "nine",
				metricName:  "New counter",
			},
			wantErr: storage.ErrorIncorrectTypeInt64,
		},
		{
			name: "Negative test #1",
			args: args{
				metricType:  "idk",
				metricValue: "123",
				metricName:  "New idk",
			},
			wantErr: storage.ErrorIncorrectTypeMetrics,
		},
	}
	storage.ClearStorage()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := UpdateMetrics(tt.args.metricType, tt.args.metricValue, tt.args.metricName)
			if err != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
			}
		})
	}
}
