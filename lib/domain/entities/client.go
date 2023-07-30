package entities

// Client example
type Client struct {
	Id     int    `json:"id,omitempty" example:"1"`
	Name   string `json:"name,omitempty" example:"Mobile Client"`
	Secret string `json:"secret,omitempty" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTEwLTIzVDEwOjU0OjUwLjQ4OTg0MiswNzowMCJ9.vl8apKYm9UQbj1qaG-BB2eStTEYy1ZJpPoVyuoXDr1k"`
}
