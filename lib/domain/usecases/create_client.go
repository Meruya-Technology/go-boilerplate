package usecases

import (
	"errors"

	_ "github.com/Meruya-Technology/go-boilerplate/docs"
	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	"github.com/Meruya-Technology/go-boilerplate/lib/common/https"
	htt "github.com/Meruya-Technology/go-boilerplate/lib/common/https"
	rep "github.com/Meruya-Technology/go-boilerplate/lib/domain/repositories"
	req "github.com/Meruya-Technology/go-boilerplate/lib/presentation/schemes/requests"
	ech "github.com/labstack/echo/v4"
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
func (r CreateClient) Execute(ctx ech.Context) error {
	/// Compile request
	request := req.CreateClientRequest{}
	err := htt.ParsingParameter(ctx, &request)
	if err != nil {
		return https.ErrorInternalServerResponse(ctx, err, nil)
	}

	/// Request validation
	err = r.validate(ctx, request)
	if err != nil {
		return https.ErrorBadRequest(ctx, err, nil)
	}

	/// Build & run usecase
	result, err := r.build(ctx, request)
	if err != nil {
		return https.ErrorInternalServerResponse(ctx, err, nil)
	}

	/// Return final result
	return https.CreatedResponse(ctx, "Client created successfuly", result)
}

func (r CreateClient) validate(ctx ech.Context, Request req.CreateClientRequest) error {
	config := cfg.ConfigHandler{}
	SecretKey := config.LoadConfig().Secret
	if Request.Secret != SecretKey {
		return errors.New("Invalid secret key")
	}
	return nil
}

func (r CreateClient) build(ctx ech.Context, Request req.CreateClientRequest) (interface{}, error) {
	repository := r.Repository
	result, err := repository.Create(ctx, Request)
	return result, err
}
