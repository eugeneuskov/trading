package structures

type Auth struct {
	apiKey    string
	apiSecret string
}

func NewAuth(apiKey string, apiSecret string) Auth {
	return Auth{
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}
}

func (a Auth) ApiKey() string {
	return a.apiKey
}

func (a Auth) ApiSecret() string {
	return a.apiSecret
}
