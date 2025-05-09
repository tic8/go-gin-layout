package config

import "github.com/spf13/viper"

var Config *viper.Viper
var Cfg *AppConfig

type AppConfig struct {
	Env      string         `json:"env" yaml:"env"`
	Debug    bool           `json:"debug" yaml:"debug"`
	LogPath  string         `json:"logPath" yaml:"logPath"`
	Server   ServerConfig   `json:"server" yaml:"server"`
	Redis    RedisConfig    `json:"redis" yaml:"redis"`
	Database DatabaseConfig `json:"database" yaml:"database"`
}

type RedisConfig struct {
	Addr     string `json:"addr" yaml:"addr"`
	Password string `json:"password" yaml:"password"`
	DB       int    `json:"db" yaml:"db"`
	Debug    bool   `json:"debug" yaml:"debug"`
}

type DatabaseConfig struct {
	DSN   string `json:"dsn" yaml:"dsn"`
	Debug bool   `json:"debug" yaml:"debug"`
}

type ServerConfig struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
}
