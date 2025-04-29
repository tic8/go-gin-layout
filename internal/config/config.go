package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

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

	Cfg = &AppConfig{}
	err = Config.Unmarshal(Cfg)
	if err != nil {
		panic(fmt.Errorf("Unable to decode into struct: %w", err))
	}
}
