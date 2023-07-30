package entities

// User example
type User struct {
	Id        int    `json:"id" example:"1"`
	Firstname string `json:"firstName" example:"Dwi Kurnianto"`
	Lastname  string `json:"lastName" example:"Mulyadien"`
	Phone     string `json:"phone" example:"082115100216"`
	Email     string `json:"email" example:"dwikurnianto.mulyadien@gmail.com"`
	Password  string `json:"password" example:"123456"`
}

type Users []User
