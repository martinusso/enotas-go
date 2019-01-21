package entity

type ConfigNFSe struct {
	Sequencial     int    `json:"sequencialNFe"`
	Serie          string `json:"serieNFe"`
	SequencialLote int    `json:"sequencialLoteNFe"`
	Usuario        string `jsson:"usuarioAcessoProvedor"`
	Senha          string `json:"senhaAcessoProvedor"`
	TokenAcesso    string `json:"tokenAcessoProvedor"`
}

type Empresa struct {
	ID                           string     `json:"id"`
	Status                       string     `json:"status"`
	DadosObrigatoriosPreenchidos bool       `json:"dadosObrigatoriosPreenchidos"`
	CNPJ                         string     `json:"cnpj"`
	InscricaoMunicipal           string     `json:"inscricaoMunicipal"`
	InscricaoEstadual            string     `json:"inscricaoEstadual"`
	RazaoSocial                  string     `json:"razaoSocial"`
	NomeFantasia                 string     `json:"nomeFantasia"`
	OptanteSimplesNacional       bool       `json:"optanteSimplesNacional"`
	Email                        string     `json:"email"`
	TelefoneComercial            string     `json:"telefoneComercial"`
	IncentivadorCultural         bool       `json:"incentivadorCultural"`
	RegimeEspecialTributacao     string     `json:"regimeEspecialTributacao"`
	CodigoServicoMunicipal       string     `json:"codigoServicoMunicipal"`
	ItemListaServicoLC116        string     `json:"itemListaServicoLC116"`
	CNAE                         string     `json:"cnae"`
	AliquotaISS                  float64    `json:"aliquotaIss"`
	DescricaoServico             string     `json:"descricaoServico"`
	EnviarEmailCliente           bool       `json:"enviarEmailCliente"`
	ConfiguracoesNFSeHomologacao ConfigNFSe `json:"configuracoesNFSeHomologacao"`
	ConfiguracoesNFSeProducao    ConfigNFSe `json:"configuracoesNFSeProducao"`
	NomeEmpresaCertificado       string     `json:"nome"`
	DataVencimentoCertificado    string     `json:"dataVencimento"`
	Endereco                     Endereco   `json:"endereco"`
}
