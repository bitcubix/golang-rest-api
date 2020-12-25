package app

import (
	"github.com/bitcubix/go-rest-api-boilerplate/pkg/db"
	"github.com/bitcubix/go-rest-api-boilerplate/pkg/log"
	"github.com/bitcubix/go-rest-api-boilerplate/pkg/router"
	"net/http"
	"os"
)

type Server struct {
	httpServer *http.Server
	Log        log.Logger
	Config     *Config
	DB         *db.Connection
	Router     *router.Mux
}

func New() (*Server, error) {
	var err error
	var server Server
	server.Config = LoadConfig()
	server.Log = log.New(os.Stderr, server.Config.Log.Level, server.Config.Log.Dir)
	server.DB, err = db.New(db.DriverMySQL, server.Config.Database.DSN(), server.Log.WithPrefix("db"))
	if err != nil {
		return nil, db.ErrDBConnectionFailed
	}

	return &server, nil
}
