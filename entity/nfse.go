package entity

import "github.com/hashicorp/go-uuid"

// NFSE representa a Nota Fiscal de Serviço Eletrônica
type NFSe struct {
	IdExterno      string   `json:"idExterno"`
	Ambiente       Ambiente `json:"ambienteEmissao"`
	EnviarPorEmail bool     `json:"enviarPorEmail"`
	ValorTotal     float64  `json:"valorTotal"`
	Cliente        Cliente  `json:"cliente"`
	Servico        Servico  `json:"servico"`
}

// NewNFSe cria um nova nota fiscal
func NewNFSe(c Cliente, s Servico, v float64, a Ambiente) *NFSe {
	return &NFSe{
		IdExterno:      newUUID(),
		Ambiente:       a,
		EnviarPorEmail: true,
		Cliente:        c,
		Servico:        s,
		ValorTotal:     v,
	}
}

func newUUID() (id string) {
	id, _ = uuid.GenerateUUID()
	return
}
