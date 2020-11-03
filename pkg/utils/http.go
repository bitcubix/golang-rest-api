package utils

import (
	"github.com/gabrielix29/go-rest-api/pkg/logger"

	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (r *Response) Send(w http.ResponseWriter) {
	w.WriteHeader(r.Status)
	bytes, _ := json.Marshal(r)
	_, err := w.Write(bytes)
	if err != nil {
		er := ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
		}
		er.Send(w)
	}
}

func (e *ErrorResponse) Send(w http.ResponseWriter) {
	w.WriteHeader(e.Status)
	bytes, _ := json.Marshal(e)
	_, err := w.Write(bytes)
	if err != nil {
		logger.Debug(err)
	}
}

func GetMuxParam(r *http.Request, index string) string {
	params := mux.Vars(r)
	return params[index]
}
