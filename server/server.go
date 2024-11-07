package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aa-ar/httpx/handler"
	"github.com/gorilla/mux"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

type Server struct {
	origins    []string
	underlying *http.Server
	router     *mux.Router
}

func NewServer(origins ...string) *Server {
	r := setupDefaultHandlers(mux.NewRouter())
	return &Server{
		origins:    origins,
		underlying: &http.Server{},
		router:     r,
	}
}

func (s *Server) AttachHandler(handler handler.Handler) *Server {
	s.router.Handle(handler.Path(), Handler(handler.Handler)).Methods(handler.Method())
	return s
}

func (s *Server) getUnderlyingHandler() http.Handler {
	m := negroni.New()
	m.Use(gzip.Gzip(gzip.DefaultCompression))
	m.Use(cors.New(cors.Options{
		AllowedOrigins: s.origins,
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodDelete,
		},
		AllowCredentials: true,
	}))
	m.UseFunc(handler.Json)
	m.UseHandler(s.router)
	return m
}

func (s *Server) Start(p int) {
	s.underlying.Addr = fmt.Sprintf(":%d", p)
	s.underlying.Handler = s.getUnderlyingHandler()
	log.Fatal(s.underlying.ListenAndServe())
}
