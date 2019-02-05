package entity

import (
	"testing"
)

func TestNewNFSe(t *testing.T) {
	end := Endereco{}
	cliente := NewCliente("Mony", "89460862039", end)
	servico := NewServico("Serviço prestado", "1234567890", "1234", 5)
	nfse := NewNFSe(cliente, servico, 100, Homologacao)

	if idExterno := nfse.IdExterno; len(idExterno) != 36 {
		t.Errorf("Expected a valid idExterno, got '%s'", idExterno)
	}

	if enviarPorEmail := nfse.EnviarPorEmail; !enviarPorEmail {
		t.Errorf("EnviarPorEmail should not be false")
	}
}
