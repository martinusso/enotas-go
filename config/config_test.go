package config

import (
	"testing"
)

func TestConfigure(t *testing.T) {
	apiKey := "1q2w3e4r5t6y7u8i9o0p"
	c := Configure(apiKey)
	if k := c.ApiKey; k != apiKey {
		t.Errorf("Expected '%s' got '%s'", apiKey, k)
	}
}
