package server

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSendReq(t *testing.T) {
	tests := []struct {
		name   string
		server *Server
	}{
		{
			name:   "Simple test #1",
			server: NewRequest("https://www.google.com/", "GET"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code := tt.server.SendReq()
			assert.Equal(t, code, 200, "Response code get %d", code)
		})
	}
}
