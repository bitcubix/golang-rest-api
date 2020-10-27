package server

import (
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
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
	log.Fatal(http.ListenAndServe(addr, s.Router))
}
