package utils

type GeneralResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type GeneralError struct {
	Status string `json:"status"`
	Error  string `json:"message"`
}

func SetGeneralResponse(status, message string, data any) GeneralResponse {
	return GeneralResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func SetErrorResponse(status, err string) GeneralError {
	return GeneralError{
		Status: status,
		Error:  err,
	}
}
