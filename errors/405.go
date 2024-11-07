package errors

import "net/http"

type MethodNotAllowedError struct{}

func (err MethodNotAllowedError) Status() int {
	return http.StatusMethodNotAllowed
}

func (err MethodNotAllowedError) Error() string {
	return "METHOD_NOT_ALLOWED"
}
