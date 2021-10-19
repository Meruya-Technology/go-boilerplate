package utils

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
)

type ParsingError struct {
	Message string
}

func (p *ParsingError) Error() string { return p.Message }

// * errorType used for choosing error types
func ErrorType(err error) (int, error) {
	switch {
	case isPqError(err):
		return PqError(err)
	}
	return commonError(err)
}

// * isPqError used to check error if error is pg error
func isPqError(err error) bool {
	if _, ok := err.(*pq.Error); ok {
		return true
	}
	return false
}

func commonError(err error) (int, error) {
	return http.StatusInternalServerError, fmt.Errorf(err.Error())
}

// * switchErrorValidation used to check error on request validation
func SwitchErrorValidation(err error) (message string) {
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for idx, err := range castedObject {
			field := ToSnakeCase(err.Field())
			switch field {

			}

			switch err.Tag() {
			case "required":
				message = fmt.Sprintf("%s is mandatory",
					field)
			default:
				message = err.Error()
			}

			if idx == 0 {
				break
			}
		}
	}
	return
}

var pqErrorMap = map[string]int{
	"unique_violation": http.StatusConflict,
}

// PqError is
func PqError(err error) (int, error) {
	re := regexp.MustCompile("\\((.*?)\\)")
	if err, ok := err.(*pq.Error); ok {
		match := re.FindStringSubmatch(err.Detail)
		// Change Field Name
		switch match[1] {
		}

		switch err.Code.Name() {
		case "unique_violation":
			return pqErrorMap["unique_violation"], fmt.Errorf("%s already exists", match[1])
		}
	}

	return http.StatusInternalServerError, fmt.Errorf(err.Error())
}

// var commonErrorMap = map[error]int{
// 	cnt.ErrorPgProductNotFound: http.StatusNotFound,
// 	cnt.ErrorCommonOutOfStock:  http.StatusBadRequest,
// }

// // CommonError is
// func CommonError(err error) (int, error) {
// 	switch err {
// 	case cnt.ErrorPgProductNotFound:
// 		return commonErrorMap[cnt.ErrorPgProductNotFound], cnt.ErrorPgProductNotFound
// 	case cnt.ErrorCommonOutOfStock:
// 		return commonErrorMap[cnt.ErrorCommonOutOfStock], cnt.ErrorCommonOutOfStock
// 	}
// 	return http.StatusInternalServerError, fmt.Errorf(err.Error())
// }
