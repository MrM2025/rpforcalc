package application_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MrM2025/rpforcalc/tree/master/calc_go/internal/application"
	"github.com/MrM2025/rpforcalc/tree/master/calc_go/pkg/calculation"
)

// структура запроса
type RequestBody struct {
	Expression string `json:"expression"`
}
// верный запрос
func TestCalcHandler_Success(t *testing.T) {
	expected := "2"	

	r := httptest.NewRequest(http.MethodGet, "/api/v1/calculate?expression=", nil)
	w := httptest.NewRecorder()
	application.CalcHandler(w, r)

	res := w.Result()

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if string(data) != expected {
		fmt.Fprintf(w, "error")
	}

	if w.Result().StatusCode != http.StatusOK {
		fmt.Fprintf(w, "StatusCode error")
	}

	//defer r.Body.Close()


	
}

// неверное выражение
func TestCalcHandler_InvalidExpression(t *testing.T) {

	handler := http.HandlerFunc(application.CalcHandler)
	server := httptest.NewServer(handler)
	defer server.Close()

	// запрос с некорректным выражением
	requestBody := RequestBody{
		Expression: "1+/",
	}
	body, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Error marshalling request body: %v", err)
	}

	// Создание POST-запроса
	req, err := http.NewRequest("POST", server.URL+"/api/v1/calculate", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Отправка запроса и получение ответа
	resp, err := server.Client().Do(req)
	if err != nil {
		t.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// проверка что статус ответа 422
	if resp.StatusCode != http.StatusUnprocessableEntity {
		t.Fatalf("Expected status 422, got %d", resp.StatusCode)
	}

	// проверка ответа
	var response application.CalcResJSON
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}
	if response.Error != calculation.IncorrectExpressionErr.Error() {
		t.Fatalf("Expected error %v, got %v", calculation.IncorrectExpressionErr, response.Error)
	}
}

// Тест ошибки деления на ноль
func TestCalcHandler_DivisionByZero(t *testing.T) {
	handler := http.HandlerFunc(application.CalcHandler)
	server := httptest.NewServer(handler)
	defer server.Close()

	// запрос с делением на ноль
	requestBody := RequestBody{
		Expression: "1/0",
	}
	body, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Error marshalling request body: %v", err)
	}

	// Создание POST-запроса
	req, err := http.NewRequest("POST", server.URL+"/api/v1/calculate", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// отправка запроса и получение ответа
	resp, err := server.Client().Do(req)
	if err != nil {
		t.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// проверка, что статус ответа 422
	if resp.StatusCode != http.StatusUnprocessableEntity {
		t.Fatalf("Expected status 422, got %d", resp.StatusCode)
	}

	// Проверка ответа
	var response application.CalcResJSON
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}
	if response.Error != calculation.DvsByZeroErr.Error() {
		t.Fatalf("Expected error %v, got %v", calculation.DvsByZeroErr, response.Error)
	}
}

// Тест на пустое выражения
func TestCalcHandler_EmptyExpression(t *testing.T) {
	handler := http.HandlerFunc(application.CalcHandler)
	server := httptest.NewServer(handler)
	defer server.Close()

	// запрос с пустым выражением
	requestBody := RequestBody{
		Expression: "",
	}
	body, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Error marshalling request body: %v", err)
	}

	// Создание POST-запроса
	req, err := http.NewRequest("POST", server.URL+"/api/v1/calculate", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// отправка запроса и получение ответа
	resp, err := server.Client().Do(req)
	if err != nil {
		t.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Проверка, что статус ответа 422
	if resp.StatusCode != http.StatusUnprocessableEntity {
		t.Fatalf("Expected status 422, got %d", resp.StatusCode)
	}

	// проверка ответа
	var response application.CalcResJSON
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}
	if response.Error != calculation.EmptyExpressionErr.Error() {
		t.Fatalf("Expected error %v, got %v", calculation.EmptyExpressionErr, response.Error)
	}
}