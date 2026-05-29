package resterr

import (
	"net/http"
)

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewRestErr(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestErr(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Bad_Request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationErr(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Bad_Request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewForbiddenErr(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Forbidden",
		Code:    http.StatusForbidden,
	}
}

func NewForbiddenValidationErr(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Forbidden",
		Code:    http.StatusForbidden,
		Causes:  causes,
	}
}

