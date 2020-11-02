package server

import (
	"github.com/gabrielix29/go-rest-api/internal/middlewares"
	"github.com/gabrielix29/go-rest-api/pkg/logger"
	"golang.org/x/net/http2"

	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"net/http"
)

type Server struct {
	Router   *mux.Router
	Database *gorm.DB
}

func New() *Server {
	var server Server
	server.Router = mux.NewRouter()
	return &server
}

func (s *Server) Run() {
	s.InitDatabase()
	s.initServices()
	s.Router.Use(middlewares.Json)
	if viper.GetBool("logger.debug") {
		s.Router.Use(middlewares.Logging)
	}
	addr := viper.GetString("server.host") + ":" + viper.GetString("server.port")
	httpserver := &http.Server{
		Addr:    addr,
		Handler: s.Router,
	}
	_ = http2.ConfigureServer(httpserver, &http2.Server{})
	logger.Info("HTTP Server started listening on ", addr)
	logger.Fatal(httpserver.ListenAndServe())
}

func (s *Server) InitDatabase() {
	var err error
	config := gorm.Config{
		Logger:      gormLogger.Default.LogMode(gormLogger.Silent),
		PrepareStmt: true,
	}
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	dbname := viper.GetString("database.name")

	switch viper.GetString("database.driver") {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)
		s.Database, err = gorm.Open(mysql.Open(dsn), &config)
		break
	case "postgres":
		dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=disable", username, password, dbname, port, host)
		s.Database, err = gorm.Open(postgres.Open(dsn), &config)
		break
	default:
		logger.Fatal("invalid database driver please use 'mysql' or 'postgresql'")
	}

	if err != nil {
		logger.Fatal(err, " lol")
	} else {
		logger.Debug("Connected to database")
	}
}
