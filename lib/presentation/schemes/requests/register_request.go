package requests

type RegisterRequest struct {
	Firstname string `json:"firstName" example:"Dwi Kurnianto"`
	Lastname  string `json:"lastName" example:"Mulyadien"`
	Phone     string `json:"phone" example:"082115100216"`
	Email     string `json:"email" example:"dwikurnianto.mulyadien@gmail.com"`
	Password  string `json:"password" example:"123456"`
}
