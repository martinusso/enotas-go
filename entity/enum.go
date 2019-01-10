package entity

type Ambiente string
type TipoPessoa string

const (
	Homologacao Ambiente = "Homologacao"
	Producao    Ambiente = "Producao"

	PessoaFisica   TipoPessoa = "F"
	PessoaJuridica TipoPessoa = "J"
)
