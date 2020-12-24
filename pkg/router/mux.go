package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

//Mux wrapper for gorilla/mux
type Mux struct {
	Router *mux.Router
}

//New return Mux with a standard mux Router
func New() *Mux {
	return &Mux{Router: mux.NewRouter()}
}

func (m *Mux) Handle(pattern string, handler http.Handler) {
	m.Router.Handle(pattern, handler)
}

func (m *Mux) Method(method, pattern string, handler http.Handler) {
	m.Router.Handle(pattern, handler).Methods(method)
}

func (m *Mux) HandleFunc(pattern string, handleFunc http.HandlerFunc) {
	m.Router.HandleFunc(pattern, handleFunc)
}

func (m *Mux) UseMiddleware(middlewares func(http.Handler) http.Handler) {
	m.Router.Use(middlewares)
}
