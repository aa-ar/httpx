package request

import "net/http"

type BadRequestBodyError struct {
	Message string `json:"message,omitempty"`
}

func (err BadRequestBodyError) Status() int {
	return http.StatusBadRequest
}

func (err BadRequestBodyError) Error() string {
	return "BAD_REQUEST_BODY"
}
