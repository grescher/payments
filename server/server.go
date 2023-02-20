package server

import (
	"net"
	"net/http"

	"github.com/zeebo/errs"
)

var serverErr = errs.Class("web server error")

type Server struct {
	listener net.Listener
	server   http.Server
}
