package entity

import (
	"encoding/json"
	"testing"
)

func TestCidadeMarshal(t *testing.T) {
	end := Endereco{
		Logradouro: "Rua",
		Numero:     "42",
		Bairro:     "Centro",
		CEP:        "12345678",
		Cidade:     "Vitória",
		UF:         "ES",
	}
	cliente := NewCliente("Mony", "656.763.380-83", end)
	j, _ := json.Marshal(cliente)
	expected := `{"tipoPessoa":"F","nome":"Mony","email":"","cpfCnpj":"65676338083","endereco":{"logradouro":"Rua","numero":"42","bairro":"Centro","cep":"12345678","cidade":"Vitória","uf":"ES"}}`
	if string(j) != expected {
		t.Errorf("Expected '%s', got '%s'", expected, string(j))
	}
}

func TestClientePessoaJuridica(t *testing.T) {
	end := Endereco{}
	cliente := NewCliente("Empresa", "33224970000120", end)
	if tipo := cliente.TipoPessoa; tipo != PessoaJuridica {
		t.Errorf("Expected '%s' got '%s'", PessoaJuridica, tipo)
	}
}
