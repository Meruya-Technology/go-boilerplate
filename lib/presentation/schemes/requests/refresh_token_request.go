package requests

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIxLTEyLTA0IDA5OjExOjEzLjYyMTQ5OSArMDcwMCBXSUIgbT0rMTA1NjA3LjQ5NzcxNTUwMiJ9.2Cs3nDjqCuHTH_TruMGFmbjIjxIg-HJetJbFbrTBBK0"`
	GrantType    string `json:"grantType" example:"Client credentials"`
	Scope        string `json:"scope" example:"*"`
}
