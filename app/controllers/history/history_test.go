package history

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHistoryInitial(t *testing.T) {

	req, err := http.NewRequest("GET", "/history", nil)

	if err != nil {

		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAll)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {

		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	respBody := strings.TrimRight(rr.Body.String(), "\r\n")
	expected := `{"history":null}`

	if respBody != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
