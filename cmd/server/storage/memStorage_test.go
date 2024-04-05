package storage

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMemStorage_GetMemStorage(t *testing.T) {
	type fields struct {
		metrics map[string]Metric
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Metric
		err    bool
	}{
		{
			name: "Positive test #1",
			fields: fields{metrics: map[string]Metric{"some_metric": {
				Gauge:   1.1,
				Counter: 1,
			}}},
			want: Metric{
				Gauge:   1.1,
				Counter: 1,
			},
			err:  false,
			args: args{key: "some_metric"},
		},
		{
			name:   "Negative test #1",
			fields: fields{metrics: make(map[string]Metric)},
			want: Metric{
				Gauge:   0,
				Counter: 12,
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage.metrics = tt.fields.metrics
			got, ok := GetMemStorage(tt.args.key)
			if ok == tt.err {
				t.Errorf("GetMemStorage() got1 = %v, want %v", ok, tt.err)
			}
			if !reflect.DeepEqual(got, tt.want) && !tt.err {
				t.Errorf("GetMemStorage() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemStorage_SetMemStorage(t *testing.T) {
	type fields struct {
		metrics map[string]Metric
	}
	type args struct {
		key   string
		value Metric
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Add only Gauge",
			fields: fields{metrics: make(map[string]Metric)},
			args: args{
				key: "Some metric",
				value: Metric{
					Gauge:   1.1,
					Counter: 0,
				},
			},
		},
		{
			name:   "Add only Counter",
			fields: fields{metrics: make(map[string]Metric)},
			args: args{
				key: "Some metric",
				value: Metric{
					Gauge:   0,
					Counter: 1,
				},
			},
		},
		{
			name:   "Replace Gauge",
			fields: fields{metrics: map[string]Metric{"Some metric": {1.1, 0}}},
			args: args{
				key: "Some metric",
				value: Metric{
					Gauge:   9.8913281,
					Counter: 0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage.metrics = tt.fields.metrics
			SetMemStorage(tt.args.key, tt.args.value)
			assert.Equal(t, storage.metrics[tt.args.key], tt.args.value)
		})
	}
}

func TestNewMemStorage(t *testing.T) {
	tests := []struct {
		name string
		want *MemStorage
	}{
		{
			name: "Simple test 1",
			want: &MemStorage{
				metrics: make(map[string]Metric),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMemStorage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMemStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}
