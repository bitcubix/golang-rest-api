package server

import (
	"fmt"
	"github.com/gabrielix29/go-rest-api/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
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
	addr := viper.GetString("server.host") + ":" + viper.GetString("server.port")
	logger.Info("HTTP Server started listening on ", addr)
	logger.Fatal(http.ListenAndServe(addr, s.Router))
}

func (s *Server) InitDatabase() {
	var err error
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")
	dbname := viper.GetString("database.name")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)
	s.Database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal(err)
	}else{
		logger.Debug("Connected to db")
	}
}
