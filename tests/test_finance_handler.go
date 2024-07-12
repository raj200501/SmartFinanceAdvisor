package tests

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "SmartFinanceAdvisor/handlers"
)

func TestGetFinanceData(t *testing.T) {
    req, err := http.NewRequest("GET", "/api/finance", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handlers.GetFinanceData)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    expected := `{"balance":5000.75,"expenses":3000.4,"income":7000}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
}
