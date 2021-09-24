package usecases

import (
	"encoding/json"
	"net/http"

	_ "github.com/Meruya-Technology/go-boilerplate/docs"
	"github.com/Meruya-Technology/go-boilerplate/lib/domain/entities"
	"github.com/Meruya-Technology/go-boilerplate/lib/domain/repositories"
	repositories_impl "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/repositories"
)

// RetrieveUser Get Profile
// @Summary Retrieve user profile
// @description Retrieve user profile
// @ID retrieve-user
// @Accept  json
// @Produce  json
// @Router /auth/profile [get]

type RetrieveUser struct{}

func (retrieveUser RetrieveUser) Execute(res http.ResponseWriter, req *http.Request) {
	repositories := new(repositories_impl.UserRepositoriesImpl)
	result := build(repositories)
	res.Header().Set("Content-type", "application/json")
	json.NewEncoder(res).Encode(result)
}

func build(r repositories.UserRepositories) entities.User {
	return r.GetProfile()
}
