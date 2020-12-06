package errors

import "net/http"

var (
	RestErrors restErrorInterface = &RestError{}
)

type restErrorInterface interface {
	NewNotFoundError(string) *RestError
	NewBadRequestError(string) *RestError
	NewInternalServerError(string) *RestError
}

type RestError struct {
	Message string `json:"message"`
	Status  int    `json:"code"`
	Error   string `json:"error"`
}

func (re *RestError) NewNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func (re *RestError) NewBadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}
func (re *RestError) NewInternalServerError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}
