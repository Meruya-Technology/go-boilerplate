package entities

type Client struct {
	Id     int    `json:"id" example:"1"`
	Secret string `json:"secret"`
}
