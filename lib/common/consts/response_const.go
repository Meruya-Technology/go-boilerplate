package consts

type ResponseStatus string

const (
	ResponseStatusSuccessText             ResponseStatus = "success"
	ResponseStatusCreatedText             ResponseStatus = "success insert data"
	ResponseStatusConflictText            ResponseStatus = "conflict"
	ResponseStatusInternalServerErrorText ResponseStatus = "internal server error"
	ResponseStatusBadRequestText          ResponseStatus = "bad request"
	ResponseStatusNotFoundText            ResponseStatus = "not found"
	ResponseStatusUnprocessableEntityText ResponseStatus = "unprocessable entity"
	ResponseStatusUnauthorized            ResponseStatus = "unauthorized"
	ResponseStatusForbidden               ResponseStatus = "forbidden"
)

type ResponseMessage string

const (
	ResponseMessageSuccessGetProducts ResponseMessage = "success get product list"
	ResponseMessageSuccessCreateOrder ResponseMessage = "success create order"
)
