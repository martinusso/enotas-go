package nfse

import (
	"fmt"

	"github.com/padmoney/enotas-go/entity"
)

type NFSe struct {
	IdEmpresa string
}

func NewNFSe(idEmpresa string) NFSe {
	return NFSe{
		IdEmpresa: idEmpresa,
	}
}

func (n NFSe) Emitir(nota entity.NFSe) {
	fmt.Println(n.IdEmpresa)
}
