package errors

import (
	"net/http"
)

type RestError struct {
	Message string        `json:"message"`
	Code    int           `json:"code"`
	Causes  []interface{} `json:"causes"`
}

func (e *RestError) Error() string {
	return e.Message
}

func NewDatabaseError() *RestError {
	return NewInternalServerError("database error", nil)
}
func NewBadRequestError(msg string) *RestError {
	return &RestError{
		Message: msg,
		Code:    http.StatusBadRequest,
		Causes:  nil,
	}
}

func NewNotFoundError(msg string) *RestError {
	return &RestError{
		Message: msg,
		Code:    http.StatusNotFound,
		Causes:  nil,
	}
}

func NewUnauthorizedError(msg string) *RestError {
	return &RestError{
		Message: msg,
		Code:    http.StatusUnauthorized,
		Causes:  nil,
	}
}

func NewInternalServerError(msg string, errs ...error) *RestError {
	var causes []interface{}
	for _, err := range errs {
		causes = append(causes, err)
	}
	return &RestError{
		Message: msg,
		Code:    http.StatusInternalServerError,
		Causes:  causes,
	}
}
