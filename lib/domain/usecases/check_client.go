package usecases

import (
	"context"

	_ "github.com/Meruya-Technology/go-boilerplate/docs"
	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	htt "github.com/Meruya-Technology/go-boilerplate/lib/common/https"
	rep "github.com/Meruya-Technology/go-boilerplate/lib/domain/repositories"
	req "github.com/Meruya-Technology/go-boilerplate/lib/presentation/schemes/requests"
	ech "github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type CheckClient struct {
	Repository rep.ClientRepository
	Config     cfg.Config
}

// CheckClient example
// @Description Check a client for authorization
// @ID client-check
// @tags Client
// @Accept  json
// @Produce  json
// @Param payload body requests.CheckClientRequest true "Request payload"
// @Success 201 {object} base.SuccessCreatedResponse{data=entities.Client} "Success response"
// @Success 500 {object} base.InternalServerError "Internal Server Error"
// @Success 400 {object} base.BadRequestError "Bad Request"
// @Success 401 {object} base.UnauthorizedError "Unauthorized"
// @Success 403 {object} base.ForbidenError "Forbiden"
// @Success 404 {object} base.NotFoundError "Not Found"
// @Router /client/check [post]
func (c CheckClient) Execute(ctx ech.Context) error {
	/// Compile request
	request := req.CheckClientRequest{}
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
	return htt.SuccessResponse(ctx, "EXSAC01001", "Client access granted", result)
}

func (c CheckClient) validate(ctx ech.Context, Request req.CheckClientRequest) error {
	return nil
}

func (c CheckClient) build(ctx context.Context, Request req.CheckClientRequest) (interface{}, error) {
	logger, _ := zap.NewProduction()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	/// Build repository
	repository := c.Repository
	result, err := repository.Check(ctx, Request)

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return result, err
}
