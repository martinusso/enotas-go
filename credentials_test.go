package enotas

import (
	"testing"
)

func TestNewCredentials(t *testing.T) {
	apiKey := "1q2w3e4r5t6y7u8i9o0p"
	c := NewCredentials(apiKey)
	if k := c.ApiKey; k != apiKey {
		t.Errorf("Expected '%s' got '%s'", apiKey, k)
	}
}
