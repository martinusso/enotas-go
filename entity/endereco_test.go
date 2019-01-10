package entity

import (
	"testing"
)

const (
	logradouro  = "Avenida Principal"
	numero      = "42"
	complemento = "Apto 601"
	bairro      = "Centro"
	cep         = "29090720"
	cidade      = "Vitória"
	uf          = "ES"
)

func TestMethodName(t *testing.T) {
	endereco := Endereco{}
	endereco.Logradouro = logradouro
	endereco.Numero = numero
	endereco.Complemento = complemento
	endereco.Bairro = bairro
	endereco.CEP = cep
	endereco.Cidade = cidade
	endereco.UF = uf

	if log := endereco.Logradouro; log != logradouro {
		t.Errorf("Expected '%s' got '%s'", logradouro, log)
	}
	if num := endereco.Numero; num != numero {
		t.Errorf("Expected '%s' got '%s'", numero, num)
	}
	if comp := endereco.Complemento; comp != complemento {
		t.Errorf("Expected '%s' got '%s'", complemento, comp)
	}
	if b := endereco.Bairro; b != bairro {
		t.Errorf("Expected '%s' got '%s'", bairro, b)
	}
	if c := endereco.CEP; c != cep {
		t.Errorf("Expected '%s' got '%s'", cep, c)
	}
	if c := endereco.Cidade; c != cidade {
		t.Errorf("Expected '%s' got '%s'", cidade, c)
	}
	if u := endereco.UF; u != uf {
		t.Errorf("Expected '%s' got '%s'", uf, u)
	}
}
