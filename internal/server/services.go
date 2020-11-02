package server

import "github.com/gabrielix29/go-rest-api/internal/api"

func (s *Server) initServices() {
	api.Init(s.Router)
}
