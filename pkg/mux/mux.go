package mux

import (
	"errors"

	"github.com/gorilla/mux"
)

var (
	// ErrMethodMismatch is returned when the method in the request does not match
	// the method defined against the route.
	ErrMethodMismatch = errors.New("method is not allowed")
	// ErrNotFound is returned when no route match is found.
	ErrNotFound = errors.New("no matching route was found")
)

// Router registers routes to be matched and dispatches a handler.
type Router struct {
	*mux.Router
}

// NewRouter returns a new router instance.
func NewRouter() *Router {
	return &Router{mux.NewRouter()}
}

// Route stores information to match a request and build URLs.
type Route struct {
	*mux.Route
}
