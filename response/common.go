package response

type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func HttpSuccessError() *BaseResponse {
	return &BaseResponse{
		Code:    200,
		Message: "Success",
	}
}

func NewClientError() *BaseResponse {
	return &BaseResponse{
		Code:    10001,
		Message: "Params error",
	}
}

func NewNotFoundError() *BaseResponse {
	return &BaseResponse{
		Code:    404,
		Message: "Not found",
	}
}

func NewServerError() *BaseResponse {
	return &BaseResponse{
		Code:    500,
		Message: "System error",
	}
}

func NewStatusError(message string) *BaseResponse {
	return &BaseResponse{
		Code:    401,
		Message: message,
	}
}

func NewAuthError(message string) *BaseResponse {
	return &BaseResponse{
		Code:    403,
		Message: message,
	}
}
