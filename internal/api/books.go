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
	api.BaseRoutes.Books.HandleFunc("", createBook).Methods("POST")
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	var book model.Book

	books, _ := book.GetAll(db)

	response := Response{
		Status:  http.StatusOK,
		Message: "books found",
		Data:    books,
	}
	response.Send(w)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book

	_ = json.NewDecoder(r.Body).Decode(&book)
	_ = book.Save(db)

	response := Response{
		Status:  http.StatusCreated,
		Message: "book created successfully",
		Data:    book,
	}
	response.Send(w)
}
