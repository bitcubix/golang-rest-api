package api

import (
	"github.com/gabrielix29/go-rest-api/pkg/model"

	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

var db *gorm.DB

func (api *API) InitBooks(database *gorm.DB) {
	db = database

	api.BaseRoutes.Books.HandleFunc("", getBooks).Methods("GET")
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	var book model.Book

	books, _ := book.GetAll(db)

	rsBytes, _ := json.Marshal(books)
	w.WriteHeader(http.StatusNotImplemented)
	w.Header().Add("content-type", "application/json")
	w.Write(rsBytes)
}
