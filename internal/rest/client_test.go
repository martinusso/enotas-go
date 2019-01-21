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
	credentials := config.NewCredentials(apiKey)

	client := NewClient(credentials)
	expected := "Basic " + apiKey
	if got := client.Header("Authorization"); got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
	if got := client.Header("nonexistent"); got != "" {
		t.Errorf("Expected no header, got '%s'", got)
	}

	// clear headers
	client.ClearHeaders()
	if got := client.Header("Authorization"); got != "" {
		t.Errorf("Expected no header, got '%s'", got)
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

	credentials := config.NewCredentials(apiKey)

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
		t.Errorf("Invalid Authorization header, got '%s'", expected)
	}
}
