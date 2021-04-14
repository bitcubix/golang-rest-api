package api

import (
	"net/http"

	"github.com/bitcubix/golang-rest-api/internal/services/health"
	"github.com/bitcubix/golang-rest-api/pkg/log"
)

type HealthEndpoint struct {
	logger  log.Logger
	service *health.Service
}

func NewHealthEndpoint(logger log.Logger, service *health.Service) *HealthEndpoint {
	return &HealthEndpoint{
		logger:  logger,
		service: service,
	}
}

func (e *HealthEndpoint) GetHealth(w http.ResponseWriter, _ *http.Request) {
	SendResponse(w, http.StatusOK, Response{"status": e.service.GetStatus()})
}
