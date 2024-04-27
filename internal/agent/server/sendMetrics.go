package server

import (
	"fmt"
	"github.com/stranik28/MetricsCollector/internal/agent/models"
	"github.com/stranik28/MetricsCollector/internal/agent/storage"
	"log"
)

func SendMetrics(memStorage *storage.MemStorage, servAddr string) {
	for _, store := range memStorage.Metrics {
		for k, v := range store.Gauge {
			model := models.Metrics{
				ID:    k,
				MType: "gauge",
				Value: &v,
			}
			url := fmt.Sprintf("http://%s/update/", servAddr)
			req := NewServer(url)
			code := req.SendReqPost("POST", model)
			if code != 200 {
				url = fmt.Sprintf("http://%s/update/gauge/%s/%f", servAddr, k, v)
				req = NewServer(url)
				code = req.SendReq("POST")
				log.Print("Ответ от сервера не 200")
			}
		}
		model := models.Metrics{
			ID:    "PollCount",
			MType: "counter",
			Delta: &store.Counter,
		}
		url := fmt.Sprintf("http://%s/update/", servAddr)
		req := NewServer(url)
		code := req.SendReqPost("POST", model)
		if code != 200 {
			url := fmt.Sprintf("http://%s/update/counter/PollCount/%d", servAddr, store.Counter)
			req = NewServer(url)
			code = req.SendReq("POST")
			log.Println("Ответ от сервера не 200")
		}
	}
	memStorage.ClearMemStorage()
}
