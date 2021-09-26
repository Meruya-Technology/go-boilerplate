package base

// BaseResponse Example
type BaseResponse struct {
	Status  int          `json:"status,omitempty"`
	Code    string       `json:"code,omitempty"`
	Message string       `json:"message,omitempty"`
	Data    *interface{} `json:"data,omitempty"`
}

// BaseResponse Example
type SuccessResponse struct {
	Status  int          `json:"status,omitempty" example:"200"`
	Code    string       `json:"code,omitempty"`
	Message string       `json:"message,omitempty" example:"Success"`
	Data    *interface{} `json:"data,omitempty"`
}

// InternalServerError Example
type InternalServerError struct {
	Status  int    `json:"status,omitempty" example:"500"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty" example:"Internal Server Error"`
}

// BadRequestError Example
type BadRequestError struct {
	Status  int    `json:"status,omitempty" example:"400"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty" example:"Bad Request"`
}

// UnauthorizedError Example
type UnauthorizedError struct {
	Status  int    `json:"status,omitempty" example:"401"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty" example:"Unauthorized"`
}

// ForbidenError Example
type ForbidenError struct {
	Status  int    `json:"status,omitempty" example:"403"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty" example:"Forbiden Access"`
}

// NotFoundError Example
type NotFoundError struct {
	Status  int    `json:"status,omitempty" example:"404"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty" example:"Not Found"`
}
