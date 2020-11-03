package api

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var db *gorm.DB

type API struct {
	BaseRoutes *Routes
}

type Routes struct {
	Root    *mux.Router
	ApiRoot *mux.Router

	Books  *mux.Router
	Author *mux.Router
}

func Init(root *mux.Router, database *gorm.DB) *API {
	var api API
	api.BaseRoutes = &Routes{}
	db = database

	api.BaseRoutes.Root = root
	api.BaseRoutes.ApiRoot = root.PathPrefix("/api/v1/").Subrouter()

	api.BaseRoutes.Books = api.BaseRoutes.ApiRoot.PathPrefix("/books").Subrouter()
	api.BaseRoutes.Author = api.BaseRoutes.ApiRoot.PathPrefix("/authors").Subrouter()

	api.InitBooks()
	api.InitAuthors()

	return &api
}
