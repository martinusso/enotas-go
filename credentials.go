package enotas

type Credentials struct {
	ApiKey string
}

func NewCredentials(apiKey string) Credentials {
	return Credentials{
		ApiKey: apiKey,
	}
}
