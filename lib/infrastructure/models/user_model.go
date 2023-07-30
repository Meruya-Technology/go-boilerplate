package models

// UserModel model
type UserModel struct {
	Id        int    `json:"id,omitempty" example:"1"`
	Firstname string `json:"firstName,omitempty" example:"Dwi Kurnianto"`
	Lastname  string `json:"lastName,omitempty" example:"Mulyadien"`
	Phone     string `json:"phone,omitempty" example:"082115100216"`
	Email     string `json:"email,omitempty" example:"dwikurnianto.mulyadien@gmail.com"`
	Password  string `json:"password,omitempty" example:"123456"`
}
