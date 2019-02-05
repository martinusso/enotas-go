package entity

type ServicoMunicipio struct {
	Codigo                            int     `json:"codigo"`
	Descricao                         string  `json:"descricao"`
	CodigoIBGECidade                  int     `json:"codigoIBGECidade"`
	AliquotaSugerida                  float64 `json:"aliquotaSugerida"`
	ConstrucaoCivil                   bool    `json:"construcaoCivil"`
	PercentualAproximadoFederalIBPT   float64 `json:"percentualAproximadoFederalIBPT"`
	PercentualAproximadoEstadualIBPT  float64 `json:"percentualAproximadoEstadualIBPT"`
	PercentualAproximadoMunicipalIBPT float64 `json:"percentualAproximadoMunicipalIBPT"`
	ChaveTabelaIBPT                   string  `json:"chaveTabelaIBPT"`
}
