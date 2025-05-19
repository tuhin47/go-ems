package config

import (
	"fmt"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

type AppConfig struct {
	Name string `json:"name"`
	Port int    `json:"port"`
}

type DbConfig struct {
	Host            string `json:"host"`
	Port            string `json:"port"`
	User            string `json:"user"`
	Pass            string `json:"pass"`
	Schema          string `json:"schema"`
	MaxIdleConn     int    `json:"maxIdleConn"`
	MaxOpenConn     int    `json:"maxOpenConn"`
	MaxConnLifetime int    `json:"maxConnLifetime"`
	Debug           bool   `json:"debug"`
}

type LoggerConfig struct {
	Level    string `json:"level"`
	FilePath string `json:"filePath"`
}

type Config struct {
	App    AppConfig    `json:"app"`
	Db     DbConfig     `json:"db"`
	Logger LoggerConfig `json:"logger"`
}

var config Config

func LoadConfig() {
	const (
		ENV_CONSUL_URL  = "CONSUL_URL"
		ENV_CONSUL_PATH = "CONSUL_PATH"
	)

	_ = viper.BindEnv(ENV_CONSUL_URL)
	_ = viper.BindEnv(ENV_CONSUL_PATH)

	consulURL := viper.GetString(ENV_CONSUL_URL)
	consulPath := viper.GetString(ENV_CONSUL_PATH)

	if len(consulURL) == 0 || len(consulPath) == 0 {
		panic("consul url or path not set")
	}

	viper.AddRemoteProvider("consul", consulURL, consulPath)
	viper.SetConfigType("json")
	if err := viper.ReadRemoteConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	fmt.Println("Config loaded from Consul")
}

func GetConfig() Config {
	return config
}

func App() *AppConfig {
	return &config.App
}

func Db() *DbConfig {
	return &config.Db
}

func Logger() *LoggerConfig {
	return &config.Logger
}
