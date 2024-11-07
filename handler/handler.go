package handler

import (
	"net/http"
)

type Handler interface {
	Path() string
	Method() string
	Handler(w http.ResponseWriter, r *http.Request) error
}
