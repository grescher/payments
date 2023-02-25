package server

import (
	"net"
	"net/http"
	"payments/server/handlers"

	"github.com/gorilla/mux"
	"github.com/zeebo/errs"
)

var serverError = errs.Class("web server error")

type Server struct {
	listener net.Listener
	server   http.Server
}

func NewServer(listener net.Listener, h *handlers.Handlers) *Server {
	router := mux.NewRouter()

	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/signup", h.Auth.SignUp).Methods(http.MethodPost)
	authRouter.HandleFunc("/login", h.Auth.Login).Methods(http.MethodPost)

	return &Server{
		listener: listener,
		server:   http.Server{Handler: router},
	}
}

func (s *Server) Run() (err error) {
	if err = s.server.Serve(s.listener); err != nil {
		return serverError.Wrap(err)
	}
	return nil
}

func (s *Server) Close() (err error) {
	if err = s.server.Close(); err != nil {
		serverError.Wrap(s.server.Close())
	}
	return nil
}
