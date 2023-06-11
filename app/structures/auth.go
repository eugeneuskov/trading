package structures

type Auth struct {
	apiKey    string
	apiSecret string
}

func NewAuth(apiKey string, apiSecret string) *Auth {
	return &Auth{apiKey, apiSecret}
}

func (a *Auth) ApiKey() string {
	return a.apiKey
}

func (a *Auth) ApiSecret() string {
	return a.apiSecret
}

type Token struct {
	token string
}

func NewToken(token string) *Token {
	return &Token{token}
}

func (t *Token) Token() string {
	return t.token
}
