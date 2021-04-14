package server

import (
	"net/http"
	"os"

	"github.com/bitcubix/golang-rest-api/internal/api"
	"github.com/bitcubix/golang-rest-api/internal/config"
	"github.com/bitcubix/golang-rest-api/internal/services"
	"github.com/bitcubix/golang-rest-api/pkg/db"
	"github.com/bitcubix/golang-rest-api/pkg/errors"
	"github.com/bitcubix/golang-rest-api/pkg/log"
	"github.com/bitcubix/golang-rest-api/pkg/mux"
)

type Server struct {
	server   *http.Server
	Log      log.Logger
	Config   *config.Config
	DB       *db.Connection
	Router   *mux.Router
	Services *services.Services
	API      *api.API
}

func New() (*Server, error) {
	config := config.Load()
	logger := log.New(os.Stderr, config.Log.Level, config.Log.File)

	database, err := db.New(
		"mysql",
		config.Database.GetDSN(),
		logger.WithPrefix("database"),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to db")
	}

	router := mux.NewRouter()

	httpServer := &http.Server{
		Handler: router,
		Addr:    config.Server.GetAddr(),
	}

	server := &Server{
		server: httpServer,
		Log:    logger,
		Config: config,
		DB:     database,
		Router: router,
	}

	server.setupServices()
	server.setupAPI()
	server.setupRouter()

	return server, nil
}

// Listen starts the http server.
func (s *Server) RunHTTP() error {
	var err error

	logger := s.Log.WithPrefix("http.server")

	logger.WithFields(log.Fields{"port": s.server.Addr}).Infof("staring http server")

	if err = s.server.ListenAndServe(); err != http.ErrServerClosed {
		logger.Errorf("failed to start http server: %v", err)
		return err
	}

	return nil
}

// TODO shutdown, migration
