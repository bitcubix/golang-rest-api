package api

import (
	"github.com/gabrielix29/go-rest-api/pkg/model"
	"github.com/gabrielix29/go-rest-api/pkg/utils"
	"strconv"

	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

var db *gorm.DB

func (api *API) InitBooks(database *gorm.DB) {
	db = database

	api.BaseRoutes.Books.HandleFunc("/", getBooks).Methods("GET")
	api.BaseRoutes.Books.HandleFunc("/", createBook).Methods("POST")
	api.BaseRoutes.Books.HandleFunc("/{id}", getBook).Methods("GET")
	api.BaseRoutes.Books.HandleFunc("/{id}", updateBook).Methods("PUT")
	api.BaseRoutes.Books.HandleFunc("/{id}", deleteBook).Methods("DELETE")
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	var book model.Book

	books, err := book.GetAll(db)
	if err != nil {
		er := utils.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "error while get data from database",
		}
		er.Send(w)
		return
	}

	if len(*books) == 0 {
		rs := utils.Response{
			Status:  http.StatusNotFound,
			Message: "no books found",
		}
		rs.Send(w)
		return
	}

	response := utils.Response{
		Status:  http.StatusOK,
		Message: "books found",
		Data:    books,
	}
	response.Send(w)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		er := utils.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid data",
		}
		er.Send(w)
		return
	}
	err = book.Save(db)
	if err != nil {
		er := utils.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		er.Send(w)
		return
	}

	response := utils.Response{
		Status:  http.StatusCreated,
		Message: "book created successfully",
		Data:    book,
	}
	response.Send(w)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book

	id, err := strconv.Atoi(utils.GetMuxParam(r, "id"))
	if err != nil {
		er := utils.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "id not valid",
		}
		er.Send(w)
		return
	}

	err = book.GetByID(db, uint(id))
	if err != nil {
		er := utils.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "error while get data from database",
		}
		er.Send(w)
		return
	}

	response := utils.Response{
		Status:  http.StatusOK,
		Message: "book found",
		Data:    book,
	}
	response.Send(w)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book

	id, err := strconv.Atoi(utils.GetMuxParam(r, "id"))
	if err != nil {
		er := utils.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "id not valid",
		}
		er.Send(w)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		er := utils.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid data",
		}
		er.Send(w)
		return
	}
	err = book.Update(db, id)
	if err != nil {
		er := utils.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		er.Send(w)
		return
	}

	response := utils.Response{
		Status:  http.StatusOK,
		Message: "book updated successfully",
		Data:    book,
	}
	response.Send(w)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book

	id, err := strconv.Atoi(utils.GetMuxParam(r, "id"))
	if err != nil {
		er := utils.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "id not valid",
		}
		er.Send(w)
		return
	}

	err = book.GetByID(db, uint(id))
	if err != nil {
		er := utils.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "error while get data from database",
		}
		er.Send(w)
		return
	}

	err = book.Delete(db, id)
	if err != nil {
		er := utils.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "error while deleting data from database",
		}
		er.Send(w)
		return
	}

	response := utils.Response{
		Status:  http.StatusOK,
		Message: "book deleted",
		Data:    book,
	}
	response.Send(w)
}
