package usecases

import (
	"encoding/json"
	"net/http"

	repo "github.com/Meruya-Technology/go-boilerplate/lib/domain/Repositories"
	"github.com/Meruya-Technology/go-boilerplate/lib/domain/entities"
)

type CreateClient struct {
	Repository repo.ClientRepository
}

// CreateClient example
// @Description An API for create new client
// @ID create-client-post
// @tags Client
// @Accept  json
// @Produce  json
// @Success 200 {object} base.SuccessResponse{data=entities.Client} "Success response"
// @Success 500 {object} base.InternalServerError "Internal Server Error"
// @Success 400 {object} base.BadRequestError "Bad Request"
// @Success 401 {object} base.UnauthorizedError "Unauthorized"
// @Success 403 {object} base.ForbidenError "Forbiden"
// @Success 404 {object} base.NotFoundError "Not Found"
// @Router /client/create [post]

func (r CreateClient) Execute(res http.ResponseWriter, req *http.Request) {
	/// Build
	result := r.build()

	/// Return
	res.Header().Set("Content-type", "application/json")
	json.NewEncoder(res).Encode(result)
}

func (r CreateClient) build() entities.Client {
	repository := r.Repository
	return repository.Create()
}
