package errors

import "net/http"

type RestError struct {
	Message string `json:"message"`
	Detail  string `json:"detail"`
	Status  int    `json:"code"`
	Error   string `json:"error"`
}

func NewBadRequestError(message, detail string) *RestError {
	return &RestError{
		Message: message,
		Detail:  detail,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(message, detail string) *RestError {
	return &RestError{
		Message: message,
		Detail:  detail,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewInternalServerError(message, detail string) *RestError {
	return &RestError{
		Message: message,
		Detail:  detail,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}


