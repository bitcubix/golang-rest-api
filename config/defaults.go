package config

import "github.com/spf13/viper"

// Set defaults config values for viper
func SetDefaults() {
	// server
	viper.SetDefault("server.host", "127.0.0.1")
	viper.SetDefault("server.port", "8080")

	// logger
	viper.SetDefault("logger.debug", false)
}
