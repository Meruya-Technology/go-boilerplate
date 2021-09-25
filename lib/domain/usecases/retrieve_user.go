package usecases

import (
	"encoding/json"
	"net/http"

	_ "github.com/Meruya-Technology/go-boilerplate/docs"
	"github.com/Meruya-Technology/go-boilerplate/lib/domain/entities"
	"github.com/Meruya-Technology/go-boilerplate/lib/domain/repositories"
	repositories_impl "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/repositories"
)

type RetrieveUser struct{}

// RetrieveUser example
// @Summary Add a new pet to the store
// @Description getProfile
// @ID get-profile
// @Accept  json
// @Produce  json
// @Param   some_id      path   int     true  "Some ID"
// @Param   some_id      body entities.User true  "Some ID"
// @Success 200 {string} string	"ok"
// @Router /test/user [get]
func (retrieveUser RetrieveUser) Execute(res http.ResponseWriter, req *http.Request) {
	repositories := new(repositories_impl.UserRepositoriesImpl)
	result := build(repositories)
	res.Header().Set("Content-type", "application/json")
	json.NewEncoder(res).Encode(result)
}

func build(r repositories.UserRepositories) entities.User {
	return r.GetProfile()
}
