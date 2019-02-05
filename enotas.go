package enotas

import (
	"github.com/padmoney/enotas-go/config"
	"github.com/padmoney/enotas-go/empresa"
	"github.com/padmoney/enotas-go/entity"
	"github.com/padmoney/enotas-go/internal/rest"
	"github.com/padmoney/enotas-go/nfse"
	"github.com/padmoney/enotas-go/servico"
)

type ENotas struct {
	credentials config.Credentials
}

func Configure(apiKey string) config.Credentials {
	return config.Configure(apiKey)
}

func NewENotas(c config.Credentials) ENotas {
	return ENotas{
		credentials: c,
	}
}

func (e ENotas) Cidades(pageNumber, pageSize int) ([]entity.Cidade, error) {
	c := e.getClient()
	return servico.NewCidadeServico(c).Consulta(pageNumber, pageSize)
}

func (e ENotas) Credentials() config.Credentials {
	return e.credentials
}

func (e ENotas) Empresa() empresa.Empresa {
	c := e.getClient()
	return empresa.NewEmpresa(c)
}

func (e ENotas) NFSe() nfse.NFSe {
	c := e.getClient()
	return nfse.NewNFSe(c)
}

func (e ENotas) ServicosMunicipais(uf, city string, pageNumber, pageSize int) ([]entity.ServicoMunicipio, error) {
	c := e.getClient()
	return servico.NewServicoMunicipio(c).Consulta(uf, city, pageNumber, pageSize)
}

func (e ENotas) getClient() rest.Client {
	return rest.NewClient(e.credentials)
}
