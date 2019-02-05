package entity

import (
	"testing"
)

const (
	descTeste        = "Descrição do serviço de teste"
	codTeste         = "1234567890"
	cnaeTeste        = "1234"
	aliquotaIssTeste = 5.0
)

func TestNewServico(t *testing.T) {
	s := NewServico(descTeste, codTeste, cnaeTeste, aliquotaIssTeste)
	if desc := s.Descricao; desc != descTeste {
		t.Errorf("Expected '%s' got '%s'", descTeste, desc)
	}
	if cod := s.CodigoServicoMunicipio; cod != codTeste {
		t.Errorf("Expected '%s' got '%s'", codTeste, cod)
	}
	if c := s.CNAE; c != cnaeTeste {
		t.Errorf("Expected '%s' got '%s'", cnaeTeste, c)
	}
	if iss := s.AliquotaISS; iss != aliquotaIssTeste {
		t.Errorf("Expected '%f' got '%f'", aliquotaIssTeste, iss)
	}
}
