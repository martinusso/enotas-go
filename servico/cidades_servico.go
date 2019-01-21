package servico

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/padmoney/enotas-go/config"
	"github.com/padmoney/enotas-go/entity"
	"github.com/padmoney/enotas-go/internal/rest"
)

type CidadeServico struct {
	client rest.Client
}

func NewCidadeServico(client rest.Client) CidadeServico {
	return CidadeServico{
		client: client,
	}
}

func (s CidadeServico) Consulta(pageNumber, pageSize int) ([]entity.Cidade, error) {
	var servicos struct {
		TotalRecords int             `json:"totalRecords"`
		Data         []entity.Cidade `json:"data"`
	}

	url := fmt.Sprintf("%s/servicos/cidades?pageNumber=%d&pageSize=%d", config.Endpoint, pageNumber, pageSize)
	response := s.client.Get(url)
	if response.Error != nil {
		return servicos.Data, response.Error
	}
	if response.Code == http.StatusBadRequest {
		return servicos.Data, response.Error
	}

	err := json.Unmarshal(response.Body, &servicos)
	return servicos.Data, err
}
