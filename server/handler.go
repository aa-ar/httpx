package server

import (
	"log/slog"
	"net/http"

	"github.com/aa-ar/httpx/errors"
	"github.com/aa-ar/httpx/handler"
	"github.com/aa-ar/httpx/response"
	"github.com/gorilla/mux"
)

type Handler func(w http.ResponseWriter, r *http.Request) (handler.Response, error)

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h != nil {
		if res, err := h(w, r); err != nil {
			e := convertError(err)
			response.New(e.Status(), e, nil).WriteTo(w)
			return
		} else if res != nil {
			res.WriteTo(w)
			return
		}
	}
}

func convertError(err error) errors.Error {
	e, ok := err.(errors.Error)
	if !ok || e.Status() == http.StatusInternalServerError {
		slog.Error(err.Error())
		return errors.InternalServerError{}
	}
	return e
}

func setupDefaultHandlers(r *mux.Router) *mux.Router {
	r.NotFoundHandler = Handler(notFoundHandler)
	r.MethodNotAllowedHandler = Handler(methodNotAllowedHandler)
	return r
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) (handler.Response, error) {
	return nil, errors.MethodNotAllowedError{}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) (handler.Response, error) {
	return nil, errors.NotFoundError{}
}
