package config

const (
	Endpoint = "https://api.enotasgw.com.br/v1"
)

func Configure(apiKey string) Credentials {
	return Credentials{ApiKey: apiKey}
}
