package nfse

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/padmoney/enotas-go/config"
	"github.com/padmoney/enotas-go/entity"
	"github.com/padmoney/enotas-go/internal/rest"
)

type NFSe struct {
	client rest.Client
}

func NewNFSe(client rest.Client) NFSe {
	return NFSe{
		client: client,
	}
}

func (n NFSe) Habilitar(empresaID string) error {
	url := fmt.Sprintf("%s/empresas/%s/habilitar", config.Endpoint, empresaID)
	response := n.client.Post(url, nil)
	if response.Error != nil {
		return response.Error
	}
	return nil
}

func (n NFSe) Cancelar(empresaID, id string) error {
	url := fmt.Sprintf("%s/empresas/%s/nfes/%s", config.Endpoint, empresaID, id)
	response := n.client.Delete(url)
	if response.Error != nil {
		return response.Error
	}
	return nil
}

func (n NFSe) Consultar(empresaID, id string) (entity.NFSe, error) {
	var nota entity.NFSe

	url := fmt.Sprintf("%s/empresas/%s/nfes/%s", config.Endpoint, empresaID, id)
	response := n.client.Get(url)
	if response.Error != nil {
		return nota, response.Error
	}
	err := json.Unmarshal(response.Body, &nota)
	return nota, err
}

func (n NFSe) Desabilitar(empresaID string) error {
	url := fmt.Sprintf("%s/empresas/%s/desabilitar", config.Endpoint, empresaID)
	response := n.client.Post(url, nil)
	if response.Error != nil {
		return response.Error
	}

	return nil
}

func (n NFSe) DownloadPDF(empresaID, id string) ([]byte, error) {
	url := fmt.Sprintf("%s/empresas/%s/nfes/%s/pdf", config.Endpoint, empresaID, id)
	response := n.client.Get(url)
	if response.Error != nil {
		var pdf []byte
		return pdf, response.Error
	}
	return response.Body, nil
}

func (n NFSe) DownloadXML(empresaID, id string) ([]byte, error) {
	url := fmt.Sprintf("%s/empresas/%s/nfes/%s/xml", config.Endpoint, empresaID, id)
	response := n.client.Get(url)
	if response.Error != nil {
		var pdf []byte
		return pdf, response.Error
	}
	return response.Body, nil
}

func (n NFSe) Emitir(empresaID string, nfse entity.NFSe) (string, error) {
	url := fmt.Sprintf("%s/empresas/%s/nfes", config.Endpoint, empresaID)
	body, err := json.Marshal(nfse)
	if err != nil {
		return "", err
	}
	response := n.client.Post(url, body)
	if response.Error != nil {
		return "", response.Error
	}
	if !response.Ok() {
		return "", errors.New("Erro ao emitir a NFSe.")
	}
	var ret struct {
		ID string `json:"nfeId"`
	}
	if err := json.Unmarshal(response.Body, &ret); err != nil {
		return "", errors.New("Erro no retorno da emissão de Nota Fiscal.")
	}
	return ret.ID, nil
}
