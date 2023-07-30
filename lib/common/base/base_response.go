package base

// BaseResponse Example
type BaseResponse struct {
	Code    *string     `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// BaseResponse Example
type SuccessResponse struct {
	Code    *string      `json:"code"`
	Message string       `json:"message" example:"Success"`
	Data    *interface{} `json:"data,omitempty"`
}

// BaseResponse Example
type SuccessCreatedResponse struct {
	Code    *string      `json:"code"`
	Message string       `json:"message" example:"Success"`
	Data    *interface{} `json:"data,omitempty"`
}

// InternalServerError Example
type InternalServerError struct {
	Code    *string `json:"code"`
	Message string  `json:"message" example:"Internal Server Error"`
}

// BadRequestError Example
type BadRequestError struct {
	Code    *string `json:"code"`
	Message string  `json:"message" example:"Bad Request"`
}

// UnauthorizedError Example
type UnauthorizedError struct {
	Code    *string `json:"code"`
	Message string  `json:"message" example:"Unauthorized"`
}

// ForbidenError Example
type ForbidenError struct {
	Code    *string `json:"code"`
	Message string  `json:"message" example:"Forbiden Access"`
}

// NotFoundError Example
type NotFoundError struct {
	Code    *string `json:"code"`
	Message string  `json:"message" example:"Not Found"`
}
