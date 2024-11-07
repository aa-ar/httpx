package errors

import (
	"net/http"
)

type BadRequestError struct {
	UnderlyingError error  `json:"-"`
	Message         string `json:"message,omitempty"`
}

func (err BadRequestError) Status() int {
	return http.StatusBadRequest
}

func (err BadRequestError) Error() string {
	if err.UnderlyingError == nil {
		return "BAD_REQUEST"
	}
	return err.UnderlyingError.Error()
}
