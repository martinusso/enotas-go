package rest

import (
	"testing"
)

func TestOk(t *testing.T) {
	response := Response{Code: 200}
	if !response.Ok() {
		t.Errorf("Expected 'OK'")
	}
}
