package consts

import (
	"fmt"
	"net/http"
)

type ErrorPg error
type ErrorCommon error

var (
	ErrorPgNoDataFound   ErrorPg = fmt.Errorf("No data found")
	ErrorPgDuplicatEntry ErrorPg = fmt.Errorf("Cannot insert duplicated data")
	ErrorPgTimeout       ErrorPg = fmt.Errorf("Failed to finish, Timeout reached")
)

var (
	ErrorInternal ErrorCommon = fmt.Errorf("Internal server error")
)

var CommonErrorMap = map[error]int{
	ErrorPgNoDataFound: http.StatusNotFound,
	ErrorInternal:      http.StatusBadRequest,
}
