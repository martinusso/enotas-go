package enotas

import "github.com/padmoney/enotas-go/nfse"

const (
	Endpoint = "https://api.enotasgw.com.br/v1"
)

type ENotas struct {
	credentials Credentials
}

func NewENotas(c Credentials) ENotas {
	return ENotas{
		credentials: c,
	}
}

func (e ENotas) Credentials() Credentials {
	return e.credentials
}

func (e ENotas) NFSe(idEmpresa string) nfse.NFSe {
	return nfse.NewNFSe(idEmpresa)
}
