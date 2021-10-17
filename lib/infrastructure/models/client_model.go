package models

// ClientModel model
type ClientModel struct {
	Id     int    `json:"id" example:"1"`
	Name   string `json:"name" example:"Web Client"`
	Secret string `json:"Secret" example:""`
}
