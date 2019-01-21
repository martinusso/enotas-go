package entity

type ServicoMunicipio struct {
	Codigo                            string  `json:"codigo"`
	Descricao                         string  `json:"descricao"`
	CodigoIBGECidade                  int     `json:"codigoIBGECidade"`
	AliquotaSugerida                  float32 `json:"aliquotaSugerida"`
	ConstrucaoCivil                   bool    `json:"construcaoCivil"`
	PercentualAproximadoFederalIBPT   float32 `json:"percentualAproximadoFederalIBPT"`
	PercentualAproximadoEstadualIBPT  float32 `json:"percentualAproximadoEstadualIBPT"`
	PercentualAproximadoMunicipalIBPT float32 `json:"percentualAproximadoMunicipalIBPT"`
	ChaveTabelaIBPT                   string  `json:"chaveTabelaIBPT"`
}
