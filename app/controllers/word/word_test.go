package word

import(
	"net/http/httptest"
	"testing"
	"net/http"
	"strings"
)


func TestTranslatePrefixG(t *testing.T) {

	body := `{"english-word": "apple"}`

	req, err := http.NewRequest("POST", "/word", strings.NewReader(body))

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
	expected := `{"gopher-word":"gapple"}`

	if respBody != expected {

		t.Errorf("handler returned unexpected body: got %v want %v",
		rr.Body.String(), expected)
	}
}

func TestTranslatePrefixXR(t *testing.T) {

	body := `{"english-word": "xray"}`

	req, err := http.NewRequest("POST", "/word", strings.NewReader(body))

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
	expected := `{"gopher-word":"gexray"}`

	if respBody != expected {

		t.Errorf("handler returned unexpected body: got %v want %v",
		rr.Body.String(), expected)
	}
}

func TestTranslateStartedConsonantAndPrefixQU(t *testing.T) {

	body := `{"english-word": "chair"}`

	req, err := http.NewRequest("POST", "/word", strings.NewReader(body))

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
	expected := `{"gopher-word":"airchogo"}`

	if respBody != expected {

		t.Errorf("handler returned unexpected body: got %v want %v",
		rr.Body.String(), expected)
	}
}

func TestTranslateWithStartedConsonant(t *testing.T) {

	body := `{"english-word": "square"}`

	req, err := http.NewRequest("POST", "/word", strings.NewReader(body))

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
	expected := `{"gopher-word":"aresquogo"}`

	if respBody != expected {

		t.Errorf("handler returned unexpected body: got %v want %v",
		rr.Body.String(), expected)
	}
}
