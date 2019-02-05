package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/padmoney/enotas-go/config"
)

const (
	apiKey = "q1w2e3r4t5y6u7i8o9p0"
)

func TestNewClient(t *testing.T) {
	credentials := config.Configure(apiKey)

	client := NewClient(credentials)
	expected := "Basic " + apiKey
	if got := client.Header("Authorization"); got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
	if got := client.Header("nonexistent"); got != "" {
		t.Errorf("Expected no header, got '%s'", got)
	}
}
func TestBadRequest(t *testing.T) {
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`[{"codigo":"002","mensagem":"Os dados obrigatórios da empresa ainda não foram preenchidos corretamente."}]`))
		}))
	defer ts.Close()

	url := ts.URL

	credentials := config.Configure(apiKey)

	client := NewClient(credentials)
	response := client.Post(url, nil)
	if response.Error == nil {
		t.Errorf("Expected '%s', got no error", response.Error.Error())
	}
	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected '400 Bad Request', got '%d'", response.Code)
	}
}

func TestDelete(t *testing.T) {
	var method string
	var headers http.Header
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			method = r.Method
			headers = r.Header
		}))
	defer ts.Close()

	url := ts.URL

	credentials := config.Configure(apiKey)

	client := NewClient(credentials)
	response := client.Delete(url)
	if response.Error != nil {
		t.Errorf("Expected no error, got '%s'", response.Error.Error())
	}
	if method != http.MethodDelete {
		t.Errorf("Expected '%s', got '%s'", http.MethodDelete, method)
	}
	expected := "Basic " + apiKey
	if headers.Get("Authorization") != expected {
		t.Errorf("Invalid Authorization header, got '%s'", expected)
	}
}

func TestGet(t *testing.T) {
	var content []byte
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			content = []byte(`{"totalRecords": 0,"data": []}`)
			w.Write(content)
		}))
	defer ts.Close()

	url := ts.URL

	credentials := config.Configure(apiKey)

	client := NewClient(credentials)
	response := client.Get(url)
	if response.Error != nil {
		t.Errorf("Expected no error, got '%s'", response.Error.Error())
	}
	if string(response.Body) != string(content) {
		t.Errorf("Expected '%s', got '%s'", content, response.Body)
	}
	expected := "Basic " + apiKey
	if client.Header("Authorization") != expected {
		t.Errorf("Invalid Authorization header, got '%s'", client.Header("Authorization"))
	}
}

func TestPost(t *testing.T) {
	var method string
	var headers http.Header
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			method = r.Method
			headers = r.Header
		}))
	defer ts.Close()

	url := ts.URL

	credentials := config.Configure(apiKey)

	client := NewClient(credentials)
	response := client.Post(url, nil)
	if response.Error != nil {
		t.Errorf("Expected no error, got '%s'", response.Error.Error())
	}
	if method != http.MethodPost {
		t.Errorf("Expected '%s', got '%s'", http.MethodPost, method)
	}
	expected := "Basic " + apiKey
	if headers.Get("Authorization") != expected {
		t.Errorf("Invalid Authorization header, got '%s'", headers.Get("Authorization"))
	}
	expected = "application/json"
	if headers.Get("Content-Type") != expected {
		t.Errorf("Invalid Content-Type header, got '%s'", headers.Get("Content-Type"))
	}
}

func TestPut(t *testing.T) {
	var method string
	var headers http.Header
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			method = r.Method
			headers = r.Header
		}))
	defer ts.Close()

	url := ts.URL

	credentials := config.Configure(apiKey)

	client := NewClient(credentials)
	response := client.Put(url, nil)
	if response.Error != nil {
		t.Errorf("Expected no error, got '%s'", response.Error.Error())
	}
	if method != http.MethodPut {
		t.Errorf("Expected '%s', got '%s'", http.MethodPut, method)
	}
	expected := "Basic " + apiKey
	if headers.Get("Authorization") != expected {
		t.Errorf("Invalid Authorization header, got '%s'", headers.Get("Authorization"))
	}
	expected = "application/json"
	if headers.Get("Content-Type") != expected {
		t.Errorf("Invalid Content-Type header, got '%s'", headers.Get("Content-Type"))
	}
}
