package authentication

type Token struct {
	JWT          string `json:"jwt"`
	RefreshToken string `json:"refresh_token"`
}

func NewToken(jwt, refreshToken string) *Token {
	return &Token{
		JWT:          jwt,
		RefreshToken: refreshToken,
	}
}
