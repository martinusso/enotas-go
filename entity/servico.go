package entity

// Servico representa o objeto contendo os dados do serviço prestado
type Servico struct {
	CNAE                   string  `json:"cnae"`
	CodigoServicoMunicipio string  `json:"codigoServicoMunicicio"`
	Descricao              string  `json:"descricao"`
	AliquotaISS            float64 `json:"aliquotaIss"`
	IssRetidoFonte         bool    `json:"issRetidoFonte"`
	ValorPIS               float64 `json:"valorPis"`
	ValorCOFINS            float64 `json:"valorCofins"`
	ValorCSLL              float64 `json:"valorCsll"`
	ValorINSS              float64 `json:"valorInss"`
	ValorIR                float64 `json:"valorIr"`
}

// NewService cria um novo serviço
func NewServico(d string, c string, cnae string, iss float64) Servico {
	return Servico{
		Descricao:              d,
		CodigoServicoMunicipio: c,
		CNAE:                   cnae,
		AliquotaISS:            iss,
	}
}
