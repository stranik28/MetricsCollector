package storage

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMemStorage_GetMemStorage(t *testing.T) {
	type fields struct {
		metrics map[string]Metrics
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Metrics
		want1  bool
	}{
		{
			name: "Simple test #1",
			fields: fields{metrics: map[string]Metrics{"some_metric": {
				Gauge:   1.1,
				Counter: 1,
			}}},
			want: Metrics{
				Gauge:   1.1,
				Counter: 1,
			},
			want1: true,
			args:  args{key: "some_metric"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mem := &MemStorage{
				metrics: tt.fields.metrics,
			}
			got, got1 := mem.GetMemStorage(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMemStorage() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetMemStorage() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMemStorage_SetMemStorage(t *testing.T) {
	type fields struct {
		metrics map[string]Metrics
	}
	type args struct {
		key   string
		value Metrics
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Simple test #1",
			fields: fields{metrics: make(map[string]Metrics)},
			args: args{
				key: "Some metric",
				value: Metrics{
					Gauge:   1.1,
					Counter: 0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mem := &MemStorage{
				metrics: tt.fields.metrics,
			}
			assert.Equal(t, 0, len(mem.metrics))
			mem.SetMemStorage(tt.args.key, tt.args.value)
			assert.Equal(t, 1, len(mem.metrics))
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
			want: NewMemStorage(),
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
