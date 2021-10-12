package usecases

import (
	"encoding/json"
	"net/http"

	_ "github.com/Meruya-Technology/go-boilerplate/docs"
	ent "github.com/Meruya-Technology/go-boilerplate/lib/domain/entities"
	rep "github.com/Meruya-Technology/go-boilerplate/lib/domain/repositories"
)

type CreateClient struct {
	Repository rep.ClientRepository
}

// CreateClient example
// @Description An API for create new client
// @ID create-client
// @tags Client
// @Accept  json
// @Produce  json
// @Success 200 {object} base.SuccessResponse{data=entities.Client} "Success response"
// @Success 500 {object} base.InternalServerError "Internal Server Error"
// @Success 400 {object} base.BadRequestError "Bad Request"
// @Success 401 {object} base.UnauthorizedError "Unauthorized"
// @Success 403 {object} base.ForbidenError "Forbiden"
// @Success 404 {object} base.NotFoundError "Not Found"
// @Router /client/create [get]
func (r CreateClient) Execute(res http.ResponseWriter, req *http.Request) {
	/// Build
	result := r.build()

	/// Return
	res.Header().Set("Content-type", "application/json")
	json.NewEncoder(res).Encode(result)
}

func (r CreateClient) build() ent.Client {
	repository := r.Repository
	return repository.Create()
}
