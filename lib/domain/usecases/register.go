package usecases

import (
	"context"

	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	htt "github.com/Meruya-Technology/go-boilerplate/lib/common/https"
	rep "github.com/Meruya-Technology/go-boilerplate/lib/domain/repositories"
	req "github.com/Meruya-Technology/go-boilerplate/lib/presentation/schemes/requests"
	ech "github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Register struct {
	Repository rep.UserRepository
	Config     cfg.Config
}

// Register example
// @Description Register
// @ID auth-register
// @tags Authentication
// @Accept  json
// @Produce  json
// @Param payload body requests.RegisterRequest true "Request payload"
// @Success 201 {object} base.SuccessCreatedResponse{data=entities.Client} "Success response"
// @Success 500 {object} base.InternalServerError "Internal Server Error"
// @Success 400 {object} base.BadRequestError "Bad Request"
// @Success 401 {object} base.UnauthorizedError "Unauthorized"
// @Success 403 {object} base.ForbidenError "Forbiden"
// @Success 404 {object} base.NotFoundError "Not Found"
// @Router /auth/register [post]
func (l Register) Execute(ctx ech.Context) error {
	/// Compile request
	request := req.RegisterRequest{}
	err := htt.ParsingParameter(ctx, &request)
	if err != nil {
		return htt.ErrorParsing(ctx, err, nil)
	}

	/// Request validation
	err = l.validate(ctx, request)
	if err != nil {
		return htt.ErrorBadRequest(ctx, err, nil)
	}

	/// Build & run usecase
	newCtx := ctx.Request().Context()
	result, err := l.build(newCtx, request)
	if err != nil {
		return htt.ErrorInternalServerResponse(ctx, err, nil)
	}

	/// Return final result
	return htt.CreatedResponse(ctx, "EXSAC06001", "User logged successfuly", result)
}

func (l Register) validate(ctx ech.Context, Request req.RegisterRequest) error {
	return nil
}

func (l Register) build(ctx context.Context, Request req.RegisterRequest) (interface{}, error) {
	logger, _ := zap.NewProduction()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	/// Build repository
	repository := l.Repository
	result, err := repository.Register(ctx, Request)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return result, err
}
