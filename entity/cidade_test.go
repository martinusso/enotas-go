package entity

import (
	"encoding/json"
	"testing"
)

func TestJSON(t *testing.T) {
	cidade := Cidade{
		CodigoIBGE:   3205309,
		CodigoIBGEUF: 32,
		UF:           "ES",
		Nome:         "Vitória",
	}
	j, err := json.Marshal(cidade)
	if err != nil {
		t.Errorf("There should not be an error, error: %s", err)
	}
	expected := `{"codigoIbge":3205309,"codigoIbgeUF":32,"uf":"ES","nome":"Vitória"}`
	if string(j) != expected {
		t.Errorf("Expected '%s', got '%s'", expected, string(j))
	}
}
