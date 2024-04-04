package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	handl "github.com/stranik28/MetricsCollector/cmd/server/handlers"
)

func TestUpdateMetrics(t *testing.T) {
	// Создаем новый экземпляр Gin для тестирования
	router := gin.Default()

	// Определяем маршрут, который будет использоваться в тесте
	router.POST("/update/:metricType/:metricName/:metricValue", handl.UpdateMetrics)

	// Создаем фейковый HTTP-запрос
	req, err := http.NewRequest("POST", "/update/gauge/some_metric/22", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// Создаем фейковый объект записи ответа
	w := httptest.NewRecorder()

	// Выполняем запрос
	router.ServeHTTP(w, req)

	// Проверяем код состояния ответа
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}

	router.GET("/upd/:metricType/:metricName/:metricValue", handl.UpdateMetrics)

	// Создаем фейковый HTTP-запрос
	req, err = http.NewRequest("POST", "/update/gauge/some_metric/22", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// Создаем фейковый объект записи ответа
	w = httptest.NewRecorder()

	// Выполняем запрос
	router.ServeHTTP(w, req)

	// Проверяем код состояния ответа
	if w.Code == http.StatusNotFound {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}

	// Другие проверки, если необходимо
}
