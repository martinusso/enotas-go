package entity

type Cidade struct {
	CodigoIBGE   int    `json:"codigoIbge"`
	CodigoIBGEUF int    `json:"codigoIbgeUF"`
	UF           string `json:"uf"`
	Nome         string `json:"nome"`
}
