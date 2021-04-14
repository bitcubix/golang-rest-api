package server

import "github.com/bitcubix/golang-rest-api/internal/api"

func (s *Server) setupAPI() {
	s.API = api.Setup(s.Log, s.Services)
	s.API.SetupRoutes(s.Router)
}
