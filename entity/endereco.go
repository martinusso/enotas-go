package entity

// Endereco do cliente
// Cidade: Nome da cidade ou seu código IBGE
// UF: Sigla do Estado (ES, MG, etc.)
type Endereco struct {
	Logradouro       string `json:"logradouro"`
	Numero           string `json:"numero"`
	Complemento      string `json:"complemento,omitempty"`
	Bairro           string `json:"bairro"`
	CEP              string `json:"cep"`
	Cidade           string `json:"cidade"`
	UF               string `json:"uf"`
	CodigoIbgeUF     int    `json:"codigoIbgeUf,omitempty"`
	CodigoIbgeCidade int    `json:"codigoIbgeCidade,omitempty"`
}
