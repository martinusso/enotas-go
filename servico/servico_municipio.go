package servico

import (
	"encoding/json"
	"fmt"

	"github.com/padmoney/enotas-go/config"
	"github.com/padmoney/enotas-go/entity"
	"github.com/padmoney/enotas-go/internal/rest"
)

type ServicoMunicipio struct {
	client rest.Client
}

func NewServicoMunicipio(client rest.Client) ServicoMunicipio {
	return ServicoMunicipio{
		client: client,
	}
}

func (s ServicoMunicipio) Consulta(uf, city string, pageNumber, pageSize int) ([]entity.ServicoMunicipio, error) {
	var servicos struct {
		TotalRecords int                       `json:"totalRecords"`
		Data         []entity.ServicoMunicipio `json:"data"`
	}

	url := fmt.Sprintf("%s/estados/%s/cidades/%s/servicos?pageNumber=%d&pageSize=%d", config.Endpoint, uf, city, pageNumber, pageSize)
	response := s.client.Get(url)
	if response.Error != nil {
		return servicos.Data, response.Error
	}

	err := json.Unmarshal(response.Body, &servicos)
	return servicos.Data, err
}
