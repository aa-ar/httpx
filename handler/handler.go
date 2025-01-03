package handler

import (
	"net/http"
)

type Response interface {
	WriteTo(w http.ResponseWriter)
}

type Handler interface {
	Path() string
	Method() string
	Handler(w http.ResponseWriter, r *http.Request) (Response, error)
}
