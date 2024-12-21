package application

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"net/http/httptest"
	"testing"
)

type TRequest struct {
	Expression string `json:"expression"`
}

type TResponse struct{
	Result string `json:"result"`
}

func TestHandler(t *testing.T) {

	tres := TResponse{Result: "4"}

	response, merr := json.Marshal(tres) // merr - Marshalling error
	if merr != nil {
		// Error
	}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/calculate?expression=3+1", nil)
	w := httptest.NewRecorder()
	CalcHandler(w, req)

	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(req.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	

	if string(data) != string(response) {
		fmt.Fprintf(w, "%s, %v", data, req)
		t.Errorf("expected %v but got %v", response, data)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("wrong status code, %v", res.StatusCode)
	}

}
