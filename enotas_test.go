package enotas

import (
	"testing"

	"github.com/padmoney/enotas-go/config"
)

func TestNewENotas(t *testing.T) {
	expected := "1234567890"

	c := config.Configure(expected)
	n := NewENotas(c)
	if apiKey := n.Credentials().ApiKey; apiKey != expected {
		t.Errorf("Expected '%s' got '%s'", expected, apiKey)
	}
}
