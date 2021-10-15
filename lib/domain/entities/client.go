package entities

// Client example
type Client struct {
	Id     int    `json:"id" example:"1"`
	Name   string `json:"name" example:"Mobile Client"`
	Secret string `json:"secret" example:"test"`
}
