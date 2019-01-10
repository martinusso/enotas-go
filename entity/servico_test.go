package entity

import (
	"testing"
)

const (
	descTeste = "Descrição do serviço de teste"
	codTeste  = "1234567890"
)

func TestNewServico(t *testing.T) {
	s := NewServico(descTeste, codTeste)
	if desc := s.Descricao; desc != descTeste {
		t.Errorf("Expected '%s' got '%s'", descTeste, desc)
	}
	if cod := s.CodigoServicoMunicipio; cod != codTeste {
		t.Errorf("Expected '%s' got '%s'", codTeste, cod)
	}
}
