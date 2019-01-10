package entity

// Servico representa o objeto contendo os dados do serviço prestado
type Servico struct {
	CNAE                   string  `json:"cnae"`
	CodigoServicoMunicipio string  `json:"codigoServicoMunicicio"`
	Descricao              string  `json:"descricao"`
	AliquotaIss            float64 `json:"aliquotaIss"`
	IssRetidoFonte         bool    `json:"issRetidoFonte"`
	ValorPis               float64 `json:"valorPis"`
	ValorCofins            float64 `json:"valorCofins"`
	ValorCsll              float64 `json:"valorCsll"`
	ValorInss              float64 `json:"valorInss"`
	ValorIr                float64 `json:"valorIr"`
}

// NewService cria um novo serviço
func NewServico(d string, c string) Servico {
	return Servico{
		Descricao:              d,
		CodigoServicoMunicipio: c,
	}
}
