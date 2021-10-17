package https

import (
	"fmt"
	"net/http"

	"github.com/Meruya-Technology/go-boilerplate/lib/common/base"
	cnt "github.com/Meruya-Technology/go-boilerplate/lib/common/consts"
	utl "github.com/Meruya-Technology/go-boilerplate/lib/common/utils"
	ech "github.com/labstack/echo/v4"
)

// SuccessResponse returns
func SuccessResponse(ctx ech.Context, message string, data interface{}) error {

	responseData := base.BaseResponse{
		Status:  http.StatusOK,
		Code:    "",
		Message: string(cnt.ResponseStatusSuccessText),
		Data:    data,
	}

	return ctx.JSON(http.StatusOK, responseData)
}

// CreatedResponse returns
func CreatedResponse(ctx ech.Context, message string, data interface{}) error {

	responseData := base.BaseResponse{
		Status:  http.StatusCreated,
		Code:    "",
		Message: message,
		Data:    data,
	}

	return ctx.JSON(http.StatusCreated, responseData)
}

// // ErrorResponse returns
// func ErrorResponse(ctx ech.Context, err error, data interface{}) error {
// 	statusCode, err := errorType(err)
// 	switch statusCode {
// 	case http.StatusConflict:
// 		return ErrorConflictResponse(ctx, err, data)
// 	case http.StatusBadRequest:
// 		return ErrorBadRequest(ctx, err, data)
// 	case http.StatusNotFound:
// 		return ErrorNotFound(ctx, err, data)
// 	case http.StatusUnauthorized:
// 		return ErrorUnauthorized(ctx, err, data)
// 	case http.StatusForbidden:
// 		return ErrorForbidden(ctx, err, data)
// 	}
// 	return ErrorInternalServerResponse(ctx, err, data)
// }
// ErrorConflictResponse returns
func ErrorConflictResponse(ctx ech.Context, err error, data interface{}) error {

	responseData := base.BaseResponse{
		Status:  http.StatusConflict,
		Code:    "",
		Message: err.Error(),
	}

	return ctx.JSON(http.StatusConflict, responseData)
}

// ErrorInternalServerResponse returns
func ErrorInternalServerResponse(ctx ech.Context, err error, data interface{}) error {

	responseData := base.BaseResponse{
		Status:  http.StatusInternalServerError,
		Code:    "",
		Message: err.Error(),
	}

	return ctx.JSON(http.StatusInternalServerError, responseData)
}

// ErrorBadRequest returns
func ErrorBadRequest(ctx ech.Context, err error, data interface{}) error {
	responseData := base.BaseResponse{
		Status:  http.StatusBadRequest,
		Code:    "",
		Message: err.Error(),
	}

	return ctx.JSON(http.StatusBadRequest, responseData)
}

// ErrorNotFound returns
func ErrorNotFound(ctx ech.Context, err error, data interface{}) error {
	responseData := base.BaseResponse{
		Status:  http.StatusNotFound,
		Code:    "",
		Message: err.Error(),
	}

	return ctx.JSON(http.StatusNotFound, responseData)
}

// ErrorParsing returns
func ErrorParsing(ctx ech.Context, err error, data interface{}) error {

	responseData := base.BaseResponse{
		Status:  http.StatusUnprocessableEntity,
		Code:    "",
		Message: err.Error(),
	}

	return ctx.JSON(http.StatusUnprocessableEntity, responseData)
}

// ErrorValidate returns
func ErrorValidate(ctx ech.Context, err error, data interface{}) error {
	message := utl.SwitchErrorValidation(err)
	responseData := base.BaseResponse{
		Status:  http.StatusBadRequest,
		Code:    "",
		Message: message,
	}

	return ctx.JSON(http.StatusBadRequest, responseData)
}

// ErrorUnauthorized returns
func ErrorUnauthorized(ctx ech.Context, err error, data interface{}) error {
	responseData := base.BaseResponse{
		Status:  http.StatusUnauthorized,
		Code:    "",
		Message: err.Error(),
	}

	return ctx.JSON(http.StatusUnauthorized, responseData)
}

// ErrorForbidden returns
func ErrorForbidden(ctx ech.Context, err error, data interface{}) error {
	responseData := base.BaseResponse{
		Status:  http.StatusForbidden,
		Code:    "",
		Message: err.Error(),
	}

	return ctx.JSON(http.StatusForbidden, responseData)
}

// ErrorDefaultResponse returns
func ErrorDefaultResponse(ctx ech.Context, statusCode int, message string) error {

	switch statusCode {
	case http.StatusConflict:
		return ErrorConflictResponse(ctx, fmt.Errorf(message), nil)
	case http.StatusBadRequest:
		return ErrorBadRequest(ctx, fmt.Errorf(message), nil)
	case http.StatusNotFound:
		return ErrorNotFound(ctx, fmt.Errorf(message), nil)
	case http.StatusUnauthorized:
		return ErrorUnauthorized(ctx, fmt.Errorf(message), nil)
	}
	return ErrorInternalServerResponse(ctx, fmt.Errorf(message), nil)
}
