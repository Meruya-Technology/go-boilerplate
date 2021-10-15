package usecases

import (
	"encoding/json"
	"net/http"

	_ "github.com/Meruya-Technology/go-boilerplate/docs"
	"github.com/Meruya-Technology/go-boilerplate/lib/common/base"
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
// @Param payload body requests.CreateClientRequest true "Client payload"
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

func (r CreateClient) build() interface{} {
	repository := r.Repository
	result, err := repository.Create()
	errorHandled := r.errorHandle(err)
	if errorHandled != nil {
		return errorHandled
	}
	return result
}

func (r CreateClient) errorHandle(err error) interface{} {
	if err != nil {
		return base.InternalServerError{}
	}
	return nil
}
