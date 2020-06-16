package errors

import (
	"net/http"
	"errors"
)

type RestErr struct {
	Message string
	Status int
	Error string
	 
}
func NewError(msg string) error{
 return errors.New(msg)
}
func NewBadRequestError(message string)*RestErr{
	return &RestErr{
		Message: message,
		Status: http.StatusBadRequest,
		Error:"Bad_Request",

	}
}
func NewNotFoundError(message string)*RestErr{
	return &RestErr{
		Message: message,
		Status: http.StatusNotFound,
		Error:"Not_Found",

	}
}

func NewInternalServerError(message string)*RestErr {
	return &RestErr{
		Message: message,
		Status: http.StatusInternalServerError,
		Error:"Internal_server_error",

	}
	
}