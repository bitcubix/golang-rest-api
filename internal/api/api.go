package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
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

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
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

func (r *Response) Send(w http.ResponseWriter) {
	w.WriteHeader(r.Status)
	bytes, _ := json.Marshal(r)
	_, _ = w.Write(bytes)
	//TODO error response
}
