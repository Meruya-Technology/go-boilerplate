package usecases

import (
	"encoding/json"
	"net/http"

	"github.com/Meruya-Technology/go-boilerplate/lib/domain/entities"
	"github.com/Meruya-Technology/go-boilerplate/lib/domain/repositories"
	repositories_impl "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/repositories"
)

type RetrieveUser struct{}

var user entities.User

func (retrieveUser RetrieveUser) Execute(res http.ResponseWriter, req *http.Request) {

	repositories := new(repositories_impl.UserRepositoriesImpl)
	result := build(repositories)

	res.Header().Set("Content-type", "application/json")
	json.NewEncoder(res).Encode(result)
}

func build(r repositories.UserRepositories) entities.User {
	return r.GetProfile()
}
