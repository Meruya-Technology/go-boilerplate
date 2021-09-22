package usecases

import (
	"encoding/json"
	"net/http"

	"github.com/Meruya-Technology/go-boilerplate/lib/domain/entities"
	"github.com/Meruya-Technology/go-boilerplate/lib/domain/repositories"
	repositories_impl "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/repositories"
)

// RetrieveUser example
// @Summary Add a new pet to the store
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param   some_id      path   int     true  "Some ID"
// @Param   some_id      body web.Pet true  "Some ID"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /testapi/get-string-by-int/{some_id} [get]

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
