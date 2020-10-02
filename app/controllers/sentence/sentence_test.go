package sentence

import(
	"bytes"
	"net/http/httptest"
	"testing"
	"net/http"
	"strings"
)


func TestTranslate(t *testing.T) {

	body := []byte(`{"english-sentence": "Apple xray chair square?"}`)

	req, err := http.NewRequest("POST", "/sentence", bytes.NewReader(body))
	
	if err != nil {
	
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Translate)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {

		t.Errorf("handler returned wrong status code: got %v want %v",
		status, http.StatusOK)
	}

	respBody := strings.TrimRight(rr.Body.String(), "\r\n")
	expected := `{"gopher-sentence":"Gapple gexray airchogo aresquogo?"}`

	if respBody != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
		rr.Body.String(), expected)
	}
}
