package config

import "github.com/bitcubix/golang-rest-api/pkg/conf"

// SetDefaults gives conf the right default values for configuration
func SetDefaults() {
	// serverConfig
	conf.SetDefault("server.host", "0.0.0.0")
	conf.SetDefault("server.port", "8080")

	// logConfig
	conf.SetDefault("log.level", "info")
	conf.SetDefault("log.file", "./ecdl.log")

	// databaseConfig
	conf.SetDefault("database.username", "dbuser")
	conf.SetDefault("database.password", "password")
	conf.SetDefault("database.name", "ecdl")
	conf.SetDefault("database.host", "mariadb")
	conf.SetDefault("database.port", 3306)
}
