package app

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/spf13/viper"
	"fmt"
)

var Config appConfig

type appConfig struct {
	ServerPort string  `mapstructure:"server_port"`
}

func (config appConfig) Validate() error {
	return validation.ValidateStruct(&config,
		validation.Field(&config.ServerPort, validation.Required),
	)
}

func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("app")
	v.SetConfigType("yaml")
	v.SetDefault("server_port", 8080)
	for _, path := range configPaths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("Failed to read the configuration file: %s", err)
	}
	if err := v.Unmarshal(&Config); err != nil {
		return err
	}
	return Config.Validate()
}