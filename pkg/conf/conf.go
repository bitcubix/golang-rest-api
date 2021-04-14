package conf

import "github.com/spf13/viper"

var conf *Config

func init() {
	conf = New()
}

type Config struct {
	*viper.Viper
}

func New() *Config {
	return &Config{Viper: viper.New()}
}

func GetString(key string) string { return conf.GetString(key) }

func GetInt(key string) int { return conf.GetInt(key) }

func AddConfigPath(in string) { conf.AddConfigPath(in) }

func SetConfigName(in string) { conf.SetConfigName(in) }

func AutomaticEnv() { conf.AutomaticEnv() }

func ReadInConfig() error { return conf.ReadInConfig() }

func ConfigFileUsed() string { return conf.ConfigFileUsed() }

func SetDefault(key string, value interface{}) { conf.SetDefault(key, value) }
