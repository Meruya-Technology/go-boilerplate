package utils

import (
	"fmt"
	"net/http"
	"regexp"

	cns "github.com/Meruya-Technology/go-boilerplate/lib/common/consts"
	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
)

type ParsingError struct {
	Message string
}

func (p *ParsingError) Error() string { return p.Message }

var pqErrorMap = map[string]int{
	"unique_violation": http.StatusConflict,
}

func ErrorType(err error) (int, error) {
	switch {
	case isPqError(err):
		return PqError(err)
	}
	return commonError(err)
}

func isPqError(err error) bool {
	if _, ok := err.(*pq.Error); ok {
		return true
	}
	return false
}

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

func commonError(err error) (int, error) {
	switch err {
	case cns.ErrorPgNoDataFound:
		return cns.CommonErrorMap[cns.ErrorPgNoDataFound], cns.ErrorPgNoDataFound
	case cns.ErrorInternal:
		return cns.CommonErrorMap[cns.ErrorInternal], cns.ErrorInternal
	default:
		return http.StatusInternalServerError, fmt.Errorf(err.Error())
	}
}

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
