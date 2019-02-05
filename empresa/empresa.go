package empresa

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

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

func (e Empresa) Salvar(emp entity.Empresa) (string, error) {
	url := fmt.Sprintf("%s/empresas", config.Endpoint)
	body, err := json.Marshal(emp)
	if err != nil {
		return "", err
	}
	response := e.client.Post(url, body)
	if response.Error != nil {
		return "", response.Error
	}
	if !response.Ok() {
		return "", errors.New("Erro ao incluir/atualizar a empresa.")
	}
	var ret struct {
		ID string `json:"empresaId"`
	}
	if err := json.Unmarshal(response.Body, &ret); err != nil {
		return "", errors.New("Erro no retorno ao incluir/atualizar a empresa.")
	}
	return ret.ID, nil
}

func (e Empresa) UploadCertificado(empresaID string, password string, cert []byte) error {
	url := fmt.Sprintf("%s/empresas/{%s}/certificadoDigital", config.Endpoint, empresaID)
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	bodyWriter.WriteField("senha", password)
	fmt.Println(password)

	filename := empresaID + ".pfx"
	fileWriter, err := bodyWriter.CreateFormFile("arquivo", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}
	r := bytes.NewReader(cert)
	_, err = io.Copy(fileWriter, r)
	if err != nil {
		return err
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	e.client.AddHeader("Content-Type", contentType)
	response := e.client.Send(http.MethodPost, url, bodyBuf)
	if response.Error != nil {
		return response.Error
	}
	if !response.Ok() {
		return errors.New("Erro ao vincular certificado.")
	}
	return nil
}
