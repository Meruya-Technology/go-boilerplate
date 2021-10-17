package consts

type ResponseCode string

const (
	ErrorPgProductNotFound ResponseCode = "success get product list"
	ErrorCommonOutOfStock  ResponseCode = "success create order"
)

// var commonErrorMap = map[error]int{
// 	constant.ErrorPgProductNotFound: http.StatusNotFound,
// 	constant.ErrorCommonOutOfStock:  http.StatusBadRequest,
// }

// // CommonError is
// func CommonError(err error) (int, error) {
// 	switch err {
// 	case constant.ErrorPgProductNotFound:
// 		return commonErrorMap[constant.ErrorPgProductNotFound], constant.ErrorPgProductNotFound
// 	case constant.ErrorCommonOutOfStock:
// 		return commonErrorMap[constant.ErrorCommonOutOfStock], constant.ErrorCommonOutOfStock
// 	}
// 	return http.StatusInternalServerError, fmt.Errorf(err.Error())
// }
