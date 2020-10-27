package v1

import "github.com/gorilla/mux"

type API struct {
	BaseRoutes *Routes
}

type Routes struct {
	Root    *mux.Router
	ApiRoot *mux.Router

	Books *mux.Router
}

func Init(root *mux.Router) *API {
	var api API
	api.BaseRoutes = &Routes{}

	api.BaseRoutes.Root = root
	api.BaseRoutes.ApiRoot = root.PathPrefix("/api/v1/").Subrouter()

	api.BaseRoutes.Books = api.BaseRoutes.ApiRoot.PathPrefix("/books").Subrouter()

	api.InitBooks()

	return &api
}
