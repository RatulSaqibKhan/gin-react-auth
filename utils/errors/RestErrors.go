package errors

import (
	"net/http"
)

type RestErr struct {
	Message string
	Status int
	Error string
}

func InternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status: http.StatusInternalServerError,
		Error: "internal server error!",
	}
}

func BadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status: http.StatusBadRequest,
		Error: "bad request!",
	}
}