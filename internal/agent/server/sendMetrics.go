package server

import (
	"fmt"
	"github.com/stranik28/MetricsCollector/internal/agent/storage"
)

func SendMetrics(memStorage *storage.MemStorage, servAddr string) {
	for _, store := range memStorage.Metrics {
		for k, v := range store.Gauge {
			url := fmt.Sprintf("http://%s/update/gauge/%s/%f", servAddr, k, v)
			req := NewRequest(url, "POST")
			code := req.SendReq()
			if code != 200 {
				panic("Ответ от сервера не 200")
			}
		}
		url := fmt.Sprintf("http://%s/update/counter/PollCount/%d", servAddr, store.Counter)
		req := NewRequest(url, "POST")
		code := req.SendReq()
		if code != 200 {
			panic("Ответ от сервера не 200")
		}
	}
	memStorage.ClearMemStorage()
}
