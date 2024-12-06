package response

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func OK() Response {
	return Response{
		Status: StatusOK,
	}
}

func Error(msg string) *Response {
	return &Response{
		Status: StatusError,
		Error:  msg,
	}
}

func ValidationError(errs error) *Response {
	var errMsgs []string

	validationErrs, ok := errs.(validator.ValidationErrors)
	if !ok {
		return Error("validation failed: unexpected error format")
	}

	for _, err := range validationErrs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is a required field", err.Field()))
		case "url":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not a valid URL", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not valid", err.Field()))
		}
	}

	return &Response{
		Status: StatusError,
		Error:  strings.Join(errMsgs, ", "),
	}
}
