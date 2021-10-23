package usecases

import (
	"context"
	"errors"

	_ "github.com/Meruya-Technology/go-boilerplate/docs"
	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	htt "github.com/Meruya-Technology/go-boilerplate/lib/common/https"
	rep "github.com/Meruya-Technology/go-boilerplate/lib/domain/repositories"
	req "github.com/Meruya-Technology/go-boilerplate/lib/presentation/schemes/requests"
	ech "github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type CreateClient struct {
	Repository rep.ClientRepository
	Config     cfg.Config
}

// CreateClient example
// @Description Create a client for authorization
// @ID client-create
// @tags Client
// @Accept  json
// @Produce  json
// @Param payload body requests.CreateClientRequest true "Request payload"
// @Success 201 {object} base.SuccessCreatedResponse{data=entities.Client} "Success response"
// @Success 500 {object} base.InternalServerError "Internal Server Error"
// @Success 400 {object} base.BadRequestError "Bad Request"
// @Success 401 {object} base.UnauthorizedError "Unauthorized"
// @Success 403 {object} base.ForbidenError "Forbiden"
// @Success 404 {object} base.NotFoundError "Not Found"
// @Router /client/create [post]
func (c CreateClient) Execute(ctx ech.Context) error {
	/// Compile request
	request := req.CreateClientRequest{}
	err := htt.ParsingParameter(ctx, &request)
	if err != nil {
		return htt.ErrorParsing(ctx, err, nil)
	}

	/// Request validation
	err = c.validate(ctx, request)
	if err != nil {
		return htt.ErrorBadRequest(ctx, err, nil)
	}

	/// Build & run usecase
	newCtx := ctx.Request().Context()
	result, err := c.build(newCtx, request)
	if err != nil {
		return htt.ErrorInternalServerResponse(ctx, err, nil)
	}

	/// Return final result
	return htt.CreatedResponse(ctx, "Client created successfuly", result)
}

func (c CreateClient) validate(ctx ech.Context, Request req.CreateClientRequest) error {
	config := c.Config
	SecretKey := config.Secret
	if Request.Secret != SecretKey {
		return errors.New("Invalid secret key")
	}
	return nil
}

func (c CreateClient) build(ctx context.Context, Request req.CreateClientRequest) (interface{}, error) {
	logger, _ := zap.NewProduction()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	/// Build repository
	repository := c.Repository
	result, err := repository.Create(ctx, Request)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return result, err
}
