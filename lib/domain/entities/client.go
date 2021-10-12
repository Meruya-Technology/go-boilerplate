package entities

// Client example
type Client struct {
	Id     int    `json:"id" example:"1"`
	Secret string `json:"secret" example:"test"`
}
