package models

// User model
type User struct {
	Id    int    `json:"id" example:"1"`
	Name  string `json:"name" example:"Jhon doe"`
	Email string `json:"email" example:"johndoe@gmail.com"`
}
