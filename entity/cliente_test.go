package entity

import (
	"testing"
)

const (
	nomeMony     = "Mony"
	emailMony    = "mony@email.com"
	telefoneMony = "27 9 8765 4321"
	cpfMony      = "65676338083"
	tipoMony     = PessoaFisica
	imMony       = "1111111111"
	ieMony       = "2222222222"
	icMony       = "sim"
)

func TestCliente(t *testing.T) {
	end := Endereco{}
	cliente := NewCliente(nomeMony, cpfMony, end)
	cliente.Email = emailMony
	cliente.Telefone = telefoneMony
	cliente.InscricaoMunicipal = imMony
	cliente.InscricaoEstadual = ieMony
	cliente.IndicadorContribuinteICMS = icMony

	if nome := cliente.Nome; nome != nomeMony {
		t.Errorf("Expected '%s' got '%s'", nomeMony, nome)
	}
	if email := cliente.Email; email != emailMony {
		t.Errorf("Expected '%s' got '%s'", emailMony, email)
	}
	if tel := cliente.Telefone; tel != telefoneMony {
		t.Errorf("Expected '%s' got '%s'", telefoneMony, tel)
	}
	if cpf := cliente.CpfCnpj; cpf != cpfMony {
		t.Errorf("Expected '%s' got '%s'", cpfMony, cpf)
	}
	if tipo := cliente.TipoPessoa; tipo != tipoMony {
		t.Errorf("Expected '%s' got '%s'", tipoMony, tipo)
	}
	if im := cliente.InscricaoMunicipal; im != imMony {
		t.Errorf("Expected '%s' got '%s'", imMony, im)
	}
	if ie := cliente.InscricaoEstadual; ie != ieMony {
		t.Errorf("Expected '%s' got '%s'", ieMony, ie)
	}
	if ic := cliente.IndicadorContribuinteICMS; ic != icMony {
		t.Errorf("Expected '%s' got '%s'", icMony, ic)
	}
}

func TestClientePessoaJuridica(t *testing.T) {
	end := Endereco{}
	cliente := NewCliente("Empresa", "33224970000120", end)
	if tipo := cliente.TipoPessoa; tipo != PessoaJuridica {
		t.Errorf("Expected '%s' got '%s'", PessoaJuridica, tipo)
	}

}
