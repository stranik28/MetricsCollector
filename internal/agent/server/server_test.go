package server

import (
	"github.com/stranik28/MetricsCollector/internal/agent/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSendReq(t *testing.T) {
	logger1, err := logger.Init("info")

	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		name   string
		server *Server
	}{
		{
			name:   "Simple test #1",
			server: NewServer("https://www.google.com/"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code := tt.server.SendReq("GET", logger1)
			assert.Equal(t, code, 200, "Response code get %d", code)
		})
	}
}
