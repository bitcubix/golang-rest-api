package api

import (
	"encoding/json"
	"net/http"

	"github.com/bitcubix/golang-rest-api/internal/services"
	"github.com/bitcubix/golang-rest-api/pkg/log"
	"github.com/bitcubix/golang-rest-api/pkg/mux"
)

type Response map[string]interface{}

type API struct {
	Health *HealthEndpoint
}

func Setup(logger log.Logger, services *services.Services) *API {
	return &API{
		Health: NewHealthEndpoint(logger, services.Health),
	}
}

func (a *API) SetupRoutes(router *mux.Router) {
	apiv1 := router.PathPrefix("/api/v1").Subrouter()

	// Health Endpoint
	apiv1.HandleFunc("/health", a.Health.GetHealth).Methods(http.MethodGet)
}

func SendResponse(w http.ResponseWriter, status int, response interface{}) {
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		jsonBytesErr, _ := json.Marshal(Response{"error": "internal server error"})
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(jsonBytesErr)
		return
	}

	w.WriteHeader(status)
	_, _ = w.Write(jsonBytes)
}
