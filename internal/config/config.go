package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

var Config *viper.Viper
var Cfg *ConfigStruct

type ConfigStruct struct {
	Redis struct {
		Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
		Password string `mapstructure:"password" json:"password" yaml:"password"`
		DB       int    `mapstructure:"db" json:"db" yaml:"db"`
	} `mapstructure:"redis" json:"redis" yaml:"redis"`

	Database struct {
		DSN string `mapstructure:"dsn" json:"dsn" yaml:"dsn"`
	} `mapstructure:"database" json:"database" yaml:"database"`

	Server struct {
		Port int `mapstructure:"port" json:"port" yaml:"port"`
	} `mapstructure:"server" json:"server" yaml:"server"`
}

func InitConfig() {
	// 定义命令行参数
	configFile := flag.String("config", "", "Path to the configuration file")
	flag.Parse()

	Config = viper.New()

	if *configFile != "" {
		// 使用指定的配置文件
		Config.SetConfigFile(*configFile)
	} else {
		// 使用默认配置文件
		Config.SetConfigName("config")
		Config.SetConfigType("yaml")
		Config.AddConfigPath(".")
	}

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
