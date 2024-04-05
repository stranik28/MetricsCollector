package service

import (
	"github.com/stranik28/MetricsCollector/cmd/server/storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := UpdateMetrics(tt.args.metricType, tt.args.metricValue, tt.args.metricName)
			if err != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
			}
		})
	}
}
