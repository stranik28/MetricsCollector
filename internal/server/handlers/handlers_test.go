package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/stranik28/MetricsCollector/internal/server/models"
	"github.com/stranik28/MetricsCollector/internal/server/storage"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func makeReq(url string, method string, data models.Metrics) (*http.Request, *httptest.ResponseRecorder) {
	// Сериализуем структуру в JSON
	jsonData, _ := json.Marshal(data)

	// Создаем запрос с указанным методом и URL
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(jsonData))

	// Добавляем необходимые заголовки, например, Content-Type
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	return req, w
}

func makeReqGet(url string, method string) (*http.Request, *httptest.ResponseRecorder) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic("Could not create request")
	}
	// Создаем фейковый объект записи ответа
	w := httptest.NewRecorder()
	return req, w
}

func TestUpdateMetricsHandler(t *testing.T) {
	router := Routers("")
	value := 123.764564253
	model := models.Metrics{
		ID:    "Some metric",
		MType: "gauge",
		Delta: nil,
		Value: &value,
	}

	req, w := makeReq("/update/", "POST", model)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status code %d but got %d", http.StatusOK, w.Code)

	var respStruct models.Metrics

	err := json.Unmarshal(w.Body.Bytes(), &respStruct)

	assert.NoError(t, err)

	assert.EqualValues(t, respStruct, model)

	value = 22
	model = models.Metrics{
		ID:    "some_metric",
		MType: "gaug",
		Delta: nil,
		Value: &value,
	}

	req, w = makeReq("/update/", "POST", model)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Expected status code %d but got %d", http.StatusBadRequest, w.Code)
}

func TestGetByName(t *testing.T) {

	for k, v := range map[string]storage.Metric{"GaugeMetric": {Gauge: 6.66, Counter: 1},
		"Counter metric": {Gauge: 893482.213914, Counter: 1}} {
		storage.SetMemStorage(k, v)
	}

	router := Routers("")

	model := models.Metrics{
		ID:    "GaugeMetric",
		MType: "gauge",
	}

	//var delta int64 = 1
	//
	//modelResp := models.Metrics{
	//	ID:    "GaugeMetric",
	//	MType: "gauge",
	//	Delta: &delta,
	//}

	req, w := makeReq("/value/", "POST", model)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status code %d but got %d", http.StatusOK, w.Code)

	//assert.Equal(t, json.Unmarshal(w.Body.Bytes(), &model), modelResp)

	model.Delta = nil
	model.ID = "GaugeMetrics"
	model.MType = "gauge"
	model.Value = nil

	req, w = makeReq("/value/", "POST", model)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code, "Expected status code %d but got %d", http.StatusNotFound, w.Code)

	model.ID = "Counter metric"
	model.Delta = nil
	model.MType = "counter"
	model.Value = nil

	req, w = makeReq("/value/", "POST", model)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status code %d but got %d", http.StatusOK, w.Code)

	err := json.Unmarshal(w.Body.Bytes(), &model)
	if err != nil {
		t.Errorf("Error parsing response value: %v", err)
	}
	assert.EqualValuesf(t, *model.Delta, 1, "Actual diff")

}

func TestGetMetricHandle(t *testing.T) {
	data := map[string]storage.Metric{"GaugeMetric": {Gauge: 6.66, Counter: 1},
		"Counter metric": {Gauge: 893482.213914, Counter: 1},
		"Lol metrics":    {Gauge: 777.777777, Counter: 69}}

	for k, v := range data {
		storage.SetMemStorage(k, v)
	}
	router := Routers("")

	req, w := makeReqGet("/", "GET")

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

func TestUpdateMetricsHandlerParams(t *testing.T) {
	router := Routers("")
	req, w := makeReqGet("/update/gauge/Some metric/123.764564253", "POST")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status code %d but got %d", http.StatusOK, w.Code)

	req, w = makeReqGet("/update/gaug/some_metric/22", "POST")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Expected status code %d but got %d", http.StatusBadRequest, w.Code)
}

func TestGetByNameParams(t *testing.T) {

	for k, v := range map[string]storage.Metric{"GaugeMetric": {Gauge: 6.66, Counter: 1},
		"Counter metric": {Gauge: 893482.213914, Counter: 1}} {
		storage.SetMemStorage(k, v)
	}

	router := Routers("")

	req, w := makeReqGet("/value/gauge/GaugeMetric", "GET")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status code %d but got %d", http.StatusOK, w.Code)

	responseValue, err := strconv.ParseFloat(w.Body.String(), 64)
	if err != nil {

		t.Errorf("Error parsing response value: %v", err)
	}
	assert.Equal(t, 6.66, responseValue)

	req, w = makeReqGet("/value/gauge/GaugeMetrics", "GET")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code, "Expected status code %d but got %d", http.StatusNotFound, w.Code)

	req, w = makeReqGet("/value/counter/Counter metric", "GET")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status code %d but got %d", http.StatusOK, w.Code)

	responseValueInt, err := strconv.ParseInt(w.Body.String(), 10, 64)
	if err != nil {
		t.Errorf("Error parsing response value: %v", err)
	}
	assert.EqualValuesf(t, responseValueInt, 1, "Actual diff")

}
