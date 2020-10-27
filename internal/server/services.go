package server

import v1 "github.com/gabrielix29/go-rest-api/internal/api/v1"

func (s *Server) initServices() {
	v1.Init(s.Router)
}
