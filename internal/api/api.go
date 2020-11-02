package api

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type API struct {
	BaseRoutes *Routes
	Database   *gorm.DB
}

type Routes struct {
	Root    *mux.Router
	ApiRoot *mux.Router

	Books *mux.Router
}

func Init(root *mux.Router, database *gorm.DB) *API {
	var api API
	api.BaseRoutes = &Routes{}
	api.Database = database

	api.BaseRoutes.Root = root
	api.BaseRoutes.ApiRoot = root.PathPrefix("/api/v1/").Subrouter()

	api.BaseRoutes.Books = api.BaseRoutes.ApiRoot.PathPrefix("/books").Subrouter()

	api.InitBooks(api.Database)

	return &api
}
