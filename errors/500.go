package errors

import "net/http"

type InternalServerError struct{}

func (err InternalServerError) Status() int {
	return http.StatusInternalServerError
}

func (err InternalServerError) Error() string {
	return "INTERNAL_SERVER_ERROR"
}
