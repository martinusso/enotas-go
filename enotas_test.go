package enotas

import (
	"testing"
)

func TestNewENotas(t *testing.T) {
	expected := "1234567890"

	c := NewCredentials(expected)
	n := NewENotas(c)
	if apiKey := n.Credentials().ApiKey; apiKey != expected {
		t.Errorf("Expected '%s' got '%s'", expected, apiKey)
	}
}
