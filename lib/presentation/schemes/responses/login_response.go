package responses

type LoginResponse struct {
	AccessToken  string `json:"accessToken" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIxLTEyLTA0IDA5OjExOjEzLjYyMTQ5OSArMDcwMCBXSUIgbT0rMTA1NjA3LjQ5NzcxNTUwMiJ9.2Cs3nDjqCuHTH_TruMGFmbjIjxIg-HJetJbFbrTBBK0"`
	RefreshToken string `json:"refreshToken" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIxLTEyLTEwIDA5OjExOjEzLjYzOTcxMyArMDcwMCBXSUIgbT0rNjI0MDA3LjUxNTkyOTc5MyJ9.G7-lMOIsjJ2Y8zfpiTKME4K1XYvSlyniS3zBMYyOL5k"`
	TokenType    string `json:"tokenType" example:"Bearer"`
	ExpiresIn    int    `json:"expiresIn" example:"123456"`
}
