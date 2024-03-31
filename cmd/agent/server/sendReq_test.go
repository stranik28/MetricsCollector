package server

import "testing"

func TestSendReq(t *testing.T) {
	type Data struct {
		url    string
		method string
	}
	tests := []struct {
		name string
		data Data
	}{
		{
			name: "Simple test #1",
			data: Data{
				url:    "https://www.google.com/",
				method: "GET",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SendReq(tt.data.url, tt.data.method)
		})
	}
}
