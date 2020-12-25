package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

//Mux wrapper for gorilla/mux
type Mux struct {
	*mux.Router
}

//New return Mux with a standard mux Router
func New() *Mux {
	return &Mux{Router: mux.NewRouter()}
}

func (m *Mux) UseMiddleware(middlewares func(http.Handler) http.Handler) {
	m.Router.Use(middlewares)
}
