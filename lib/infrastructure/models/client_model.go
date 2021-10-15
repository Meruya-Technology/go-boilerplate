package models

// ClientModel model
type ClientModel struct {
	Id     int    `json:"id" example:"1"`
	Secret string `json:"secret"`
}
