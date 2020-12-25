package app

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Server   serverConfig
	Database dbConfig
	Log      logConfig
}

type serverConfig struct {
	Host string
	Port int
}

type dbConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

type logConfig struct {
	Level string
	Dir   string
}

func LoadConfig() *Config {
	setDefaults()
	loadConfig()

	return &Config{
		Server: serverConfig{
			Host: viper.GetString("server.host"),
			Port: viper.GetInt("server.port"),
		},
		Log: logConfig{
			Level: viper.GetString("log.level"),
			Dir:   viper.GetString("log.dir"),
		},
	}
}

func setDefaults() {
	viper.SetDefault("server.host", "0.0.0.0")
	viper.SetDefault("server.port", 8080)

	viper.SetDefault("database.username", "dbuser")
	viper.SetDefault("database.password", "password")
	viper.SetDefault("database.host", "127.0.0.1")
	viper.SetDefault("database.port", "3306")
	viper.SetDefault("database.name", "gorestapiboilerplate")

	viper.SetDefault("log.level", "debug")
	viper.SetDefault("log.dir", ".")
}

func loadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to load config: %s", err))
	}

	viper.WatchConfig()
	viper.SetEnvPrefix("CS")
}

func (config *dbConfig) DSN() string {
	return fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?multiStatements=true&parseTime=true&loc=UTC&collation=utf8mb4_general_ci",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)
}
