package handlers

import (
	"encoding/json"
	"github.com/stranik28/MetricsCollector/cmd/server/storage"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func makeReq(url string, method string) (*http.Request, *httptest.ResponseRecorder) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic("Could not create request")
	}
	// Создаем фейковый объект записи ответа
	w := httptest.NewRecorder()
	return req, w
}

func TestUpdateMetricsHandler(t *testing.T) {
	router := Routers()
	req, w := makeReq("/update/gauge/Some metric/123.764564253", "POST")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status code %d but got %d", http.StatusOK, w.Code)

	req, w = makeReq("/update/gaug/some_metric/22", "POST")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Expected status code %d but got %d", http.StatusBadRequest, w.Code)
}

func TestGetByName(t *testing.T) {

	for k, v := range map[string]storage.Metric{"GaugeMetric": {Gauge: 6.66, Counter: 1},
		"Counter metric": {Gauge: 893482.213914, Counter: 1}} {
		storage.SetMemStorage(k, v)
	}

	router := Routers()

	req, w := makeReq("/value/gauge/GaugeMetric", "GET")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status code %d but got %d", http.StatusOK, w.Code)

	responseValue, err := strconv.ParseFloat(w.Body.String(), 64)
	if err != nil {
		t.Errorf("Error parsing response value: %v", err)
	}
	assert.Equal(t, responseValue, 6.66)

	req, w = makeReq("/value/gauge/GaugeMetrics", "GET")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code, "Expected status code %d but got %d", http.StatusNotFound, w.Code)

	req, w = makeReq("/value/counter/Counter metric", "GET")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status code %d but got %d", http.StatusOK, w.Code)

	responseValueInt, err := strconv.ParseInt(w.Body.String(), 10, 64)
	if err != nil {
		t.Errorf("Error parsing response value: %v", err)
	}
	assert.EqualValuesf(t, responseValueInt, 1, "Actual diff")

}

func TestGetMetricHandle(t *testing.T) {
	data := map[string]storage.Metric{"GaugeMetric": {Gauge: 6.66, Counter: 1},
		"Counter metric": {Gauge: 893482.213914, Counter: 1},
		"Lol metrics":    {Gauge: 777.777777, Counter: 69}}

	for k, v := range data {
		storage.SetMemStorage(k, v)
	}
	router := Routers()

	req, w := makeReq("/", "GET")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status code %d but got %d", http.StatusOK, w.Code)

	// Декодируем ответ и сравниваем его с ожидаемыми данными
	var response map[string]storage.Metric
	err := json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Errorf("Error decoding response: %v", err)
	}

	// Проверяем, что ответ содержит все ожидаемые метрики и их значения совпадают с ожидаемыми
	for key, expectedMetric := range data {
		actualMetric, ok := response[key]
		if !ok {
			t.Errorf("Expected metric '%s' not found in response", key)
		}
		assert.Equal(t, expectedMetric, actualMetric, "Expected metric %s with value %v but got %v", key, expectedMetric, actualMetric)
	}

}
