package responses

type LoginResponse struct {
	AccessToken  string `json:"accessToken" example:"testToken"`
	RefreshToken string `json:"refreshToken" example:"textRefreshToken"`
	TokenType    string `json:"tokenType" example:"Bearer"`
	ExpiresIn    int    `json:"expiresIn" example:"123456"`
}
