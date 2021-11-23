package requests

type LoginRequest struct {
	Email    string `json:"email" example:"dwikurnaito.mulyadien@gmail.com"`
	Password string `json:"password" example:"123456"`
}
