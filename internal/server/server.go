package server

import (
	"github.com/gabrielix29/go-rest-api/pkg/logger"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"net/http"
)

type Server struct {
	Router *mux.Router
}

func New() *Server {
	var server Server
	server.Router = mux.NewRouter()
	return &server
}

func (s *Server) Run() {
	s.initServices()
	addr := viper.GetString("server.host") + ":" + viper.GetString("server.port")
	logger.Info("HTTP Server started")
	logger.Fatal(http.ListenAndServe(addr, s.Router))
}
