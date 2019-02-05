package entity

import "regexp"

type Cliente struct {
	TipoPessoa                TipoPessoa `json:"tipoPessoa"`
	Nome                      string     `json:"nome"`
	Email                     string     `json:"email"`
	CpfCnpj                   string     `json:"cpfCnpj"`
	InscricaoMunicipal        string     `json:"instricaoMunicipal,omitempty"`
	InscricaoEstadual         string     `json:"inscricaoEstadual,omitempty"`
	IndicadorContribuinteICMS string     `json:"indicadorContribuinteICMS,omitempty"`
	Telefone                  string     `json:"telefone,omitempty"`
	Endereco                  Endereco   `json:"endereco"`
}

func NewCliente(n string, d string, end Endereco) Cliente {
	c := Cliente{
		Nome:     n,
		CpfCnpj:  onlyNumbers(d),
		Endereco: end,
	}
	c.AutoTipoPessoa()
	return c
}

func (c *Cliente) AutoTipoPessoa() {
	switch len(onlyNumbers(c.CpfCnpj)) {
	case 11:
		c.TipoPessoa = PessoaFisica
	case 14:
		c.TipoPessoa = PessoaJuridica
	}
}

func onlyNumbers(s string) string {
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		return ""
	}
	return reg.ReplaceAllString(s, "")
}
