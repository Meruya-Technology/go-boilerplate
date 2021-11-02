package requests

type LoginRequest struct {
	Username string `json:"username" example:"dwikurnaitom"`
	Password string `json:"password" example:"123456"`
}
