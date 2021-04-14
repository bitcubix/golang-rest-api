package config

import (
	"github.com/bitcubix/golang-rest-api/pkg/conf"
	"github.com/bitcubix/golang-rest-api/pkg/log"
)

type Config struct {
	ConfigFile string
	Server     serverConfig
	Log        logConfig
	Database   databaseConfig
}

type serverConfig struct {
	Host string
	Port int
}

type logConfig struct {
	Level log.Level
	File  string
}

type databaseConfig struct {
	Username string
	Password string
	Name     string
	Host     string
	Port     int
}

func Load() *Config {
	logLvl, err := log.ParseLevel(conf.GetString("log.level"))
	if err != nil {
		panic(err)
	}

	return &Config{
		Server: serverConfig{
			Host: conf.GetString("server.host"),
			Port: conf.GetInt("server.port"),
		},
		Log: logConfig{
			Level: logLvl,
			File:  conf.GetString("log.file"),
		},
		Database: databaseConfig{
			Username: conf.GetString("database.username"),
			Password: conf.GetString("database.password"),
			Name:     conf.GetString("database.name"),
			Host:     conf.GetString("database.host"),
			Port:     conf.GetInt("database.port"),
		},
	}
}

// LoadFromViper loads config from the right config file
func LoadFromViper(file string) string {
	conf.AddConfigPath("./")
	conf.SetConfigName("config")

	conf.AutomaticEnv()

	if err := conf.ReadInConfig(); err == nil {
		return conf.ConfigFileUsed()
	}

	panic("no config file found use 'config -g' to generate one from defaults")
}
