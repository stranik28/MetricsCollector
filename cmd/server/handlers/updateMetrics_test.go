package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateMetrics(t *testing.T) {
	// Создаем фейковый запрос
	req := httptest.NewRequest("POST", "/update/gauge/some_metric/22", nil)
	// Создаем фейковый ответ
	w := httptest.NewRecorder()

	// Вызываем ваш обработчик с фейковым запросом и ответом
	UpdateMetrics(w, req)

	// Проверяем код статуса ответа
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}
