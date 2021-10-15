package base

// BaseResponse Example
type BaseResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    interface{}
}

// BaseResponse Example
type SuccessResponse struct {
	Status  int          `json:"status" example:"200"`
	Code    string       `json:"code"`
	Message string       `json:"message" example:"Success"`
	Data    *interface{} `json:"data"`
}

// InternalServerError Example
type InternalServerError struct {
	Status  int    `json:"status" example:"500"`
	Code    string `json:"code"`
	Message string `json:"message" example:"Internal Server Error"`
}

// BadRequestError Example
type BadRequestError struct {
	Status  int    `json:"status" example:"400"`
	Code    string `json:"code"`
	Message string `json:"message" example:"Bad Request"`
}

// UnauthorizedError Example
type UnauthorizedError struct {
	Status  int    `json:"status" example:"401"`
	Code    string `json:"code"`
	Message string `json:"message" example:"Unauthorized"`
}

// ForbidenError Example
type ForbidenError struct {
	Status  int    `json:"status" example:"403"`
	Code    string `json:"code"`
	Message string `json:"message" example:"Forbiden Access"`
}

// NotFoundError Example
type NotFoundError struct {
	Status  int    `json:"status" example:"404"`
	Code    string `json:"code"`
	Message string `json:"message" example:"Not Found"`
}
