package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var Config *viper.Viper
var Cfg *ConfigStruct

type ConfigStruct struct {
	Redis struct {
		Addr     string `mapstructure:"addr"`
		Password string `mapstructure:"password"`
		DB       int    `mapstructure:"db"`
	} `mapstructure:"redis"`

	Database struct {
		DSN string `mapstructure:"dsn"`
	} `mapstructure:"database"`

	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
}

func InitConfig() {
	Config = viper.New()
	Config.SetConfigName("config")
	Config.SetConfigType("yaml")
	Config.AddConfigPath(".")

	err := Config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w", err))
	}
	Cfg = &ConfigStruct{}
	err = Config.Unmarshal(Cfg)
	if err != nil {
		panic(fmt.Errorf("Unable to decode into struct: %w", err))
	}
}
