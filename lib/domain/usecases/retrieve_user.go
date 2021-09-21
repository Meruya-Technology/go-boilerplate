package usecases

import (
	"encoding/json"
	"net/http"

	"github.com/Meruya-Technology/go-boilerplate/lib/domain/entities"
)

type RetrieveUser struct{}

var user entities.User

func (retrieveUser RetrieveUser) Execute(res http.ResponseWriter, req *http.Request) {

	user = entities.User{
		Id:    1,
		Name:  "Jhon",
		Email: "jhondoe@gmail.com",
	}

	res.Header().Set("Content-type", "application/json")
	json.NewEncoder(res).Encode(user)
}
