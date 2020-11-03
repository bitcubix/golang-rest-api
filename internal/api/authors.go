package api

import (
	"encoding/json"
	"github.com/gabrielix29/go-rest-api/pkg/model"
	"github.com/gabrielix29/go-rest-api/pkg/utils"
	"net/http"
	"strconv"
)

func (api *API) InitAuthors() {
	api.BaseRoutes.Author.HandleFunc("/", getAuthors).Methods("GET")
	api.BaseRoutes.Author.HandleFunc("/", createAuthor).Methods("POST")
	api.BaseRoutes.Author.HandleFunc("/{id}", getAuthor).Methods("GET")
	api.BaseRoutes.Author.HandleFunc("/{id}", updateAuthor).Methods("PUT")
	api.BaseRoutes.Author.HandleFunc("/{id}", deleteAuthor).Methods("POST")
}

func getAuthors(w http.ResponseWriter, r *http.Request) {
	var author model.Author

	authors, err := author.GetAll(db)
	if err != nil {
		er := utils.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "error while get data from database",
		}
		er.Send(w)
		return
	}

	if len(*authors) == 0 {
		rs := utils.Response{
			Status:  http.StatusNotFound,
			Message: "no authors found",
		}
		rs.Send(w)
		return
	}

	response := utils.Response{
		Status:  http.StatusOK,
		Message: "authors found",
		Data:    authors,
	}
	response.Send(w)
}

func createAuthor(w http.ResponseWriter, r *http.Request) {
	var author model.Author

	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		er := utils.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid data",
		}
		er.Send(w)
		return
	}
	err = author.Save(db)
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
		Message: "author created successfully",
		Data:    author,
	}
	response.Send(w)
}

func getAuthor(w http.ResponseWriter, r *http.Request) {
	var author model.Author

	id, err := strconv.Atoi(utils.GetMuxParam(r, "id"))
	if err != nil {
		er := utils.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "id not valid",
		}
		er.Send(w)
		return
	}

	err = author.GetByID(db, uint(id))
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
		Message: "author found",
		Data:    author,
	}
	response.Send(w)
}

func updateAuthor(w http.ResponseWriter, r *http.Request) {
	var author model.Author

	id, err := strconv.Atoi(utils.GetMuxParam(r, "id"))
	if err != nil {
		er := utils.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "id not valid",
		}
		er.Send(w)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		er := utils.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid data",
		}
		er.Send(w)
		return
	}
	err = author.Update(db, id)
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
		Message: "author updated successfully",
		Data:    author,
	}
	response.Send(w)
}

func deleteAuthor(w http.ResponseWriter, r *http.Request) {
	var author model.Author

	id, err := strconv.Atoi(utils.GetMuxParam(r, "id"))
	if err != nil {
		er := utils.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "id not valid",
		}
		er.Send(w)
		return
	}

	err = author.GetByID(db, uint(id))
	if err != nil {
		er := utils.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "error while get data from database",
		}
		er.Send(w)
		return
	}

	err = author.Delete(db, id)
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
		Message: "author deleted",
		Data:    author,
	}
	response.Send(w)
}
