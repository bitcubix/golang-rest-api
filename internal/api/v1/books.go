package v1

import (
	"encoding/json"
	"github.com/gabrielix29/go-rest-api/pkg/model"
	"net/http"
)

func (api *API) InitBooks() {
	api.BaseRoutes.Books.HandleFunc("", createBook).Methods("POST")
	api.BaseRoutes.Books.HandleFunc("", getBooks).Methods("GET")
}

func createBook(w http.ResponseWriter, r *http.Request) {
	rs := model.JSONResponseMin{
		Status:  http.StatusNotImplemented,
		Message: "Book not implemented",
	}

	rsBytes, _ := json.Marshal(&rs)
	w.WriteHeader(http.StatusNotImplemented)
	w.Header().Add("content-type", "application/json")
	w.Write(rsBytes)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	rs := model.JSONResponseMin{
		Status:  http.StatusNotImplemented,
		Message: "Book not implemented",
	}

	rsBytes, _ := json.Marshal(&rs)
	w.WriteHeader(http.StatusNotImplemented)
	w.Header().Add("content-type", "application/json")
	w.Write(rsBytes)
}
