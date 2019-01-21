package empresa

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/padmoney/enotas-go/config"
	"github.com/padmoney/enotas-go/entity"
	"github.com/padmoney/enotas-go/internal/rest"
)

type Empresa struct {
	client rest.Client
}

func NewEmpresa(client rest.Client) Empresa {
	return Empresa{
		client: client,
	}
}

func (e Empresa) Consultar(id string) (entity.Empresa, error) {
	var empresa entity.Empresa

	url := fmt.Sprintf("%s/empresas/%s", config.Endpoint, id)
	response := e.client.Get(url)
	if response.Error != nil {
		return empresa, response.Error
	}
	err := json.Unmarshal(response.Body, &empresa)
	return empresa, err
}

func (e Empresa) Listar(pageNumber, pageSize int) ([]entity.Empresa, error) {
	var empresas struct {
		TotalRecords int              `json:"totalRecords"`
		Data         []entity.Empresa `json:"data"`
	}

	url := fmt.Sprintf("%s/empresas?pageNumber=%d&pageSize=%d", config.Endpoint, pageNumber, pageSize)
	response := e.client.Get(url)
	if response.Error != nil {
		return empresas.Data, response.Error
	}
	err := json.Unmarshal(response.Body, &empresas)
	return empresas.Data, err
}

func (e Empresa) Salvar(emp entity.Empresa) error {
	url := fmt.Sprintf("%s/empresas", config.Endpoint)
	body, err := json.Marshal(emp)
	if err != nil {
		return err
	}

	response := e.client.Post(url, body)
	if response.Error != nil {
		return response.Error
	}
	if response.Ok() {
		return errors.New("Erro ao incluir/atualizar a empresa.")
	}
	return nil
}
