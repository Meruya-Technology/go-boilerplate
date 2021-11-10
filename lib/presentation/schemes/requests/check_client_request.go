package requests

type CheckClientRequest struct {
	Id     int    `json:"Id" example:"1"`
	Secret string `json:"secret" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTEwLTIzVDIyOjUxOjEyLjM3NTYwOCswNzowMCJ9.TSIszUF7qDyI3EZM1NtNKVD0zZkwlbwvXfoO5uWhd9k"`
}
