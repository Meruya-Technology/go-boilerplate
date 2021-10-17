package https

import (
	utl "github.com/Meruya-Technology/go-boilerplate/lib/common/utils"
	"github.com/go-playground/validator/v10"
	ech "github.com/labstack/echo/v4"
)

// * ParsingParameter will parsing request to struct
func ParsingParameter(ctx ech.Context, i interface{}) error {
	err := ctx.Bind(i)
	if err != nil {
		return &utl.ParsingError{Message: err.Error()}
	}
	return err
}

// * ValidateParameter will validate request
func ValidateParameter(ctx ech.Context, i interface{}) (err error) {
	validate := validator.New()
	err = validate.Struct(i)
	return err
}
