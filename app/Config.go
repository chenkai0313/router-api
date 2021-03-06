package app

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/spf13/viper"
	"fmt"
)

var Config appConfig

type appConfig struct {
	ServerPort string `mapstructure:"server_port"`
	Hostname   string `mapstructure:"redis_hostname"`
	Database   int    `mapstructure:"redis_database"`
	Port       string `mapstructure:"redis_port"`
	Password   string `mapstructure:"redis_password"`
	DbDsn      string `mapstructure:"db_data_dsn"`
	Token      string `mapstructure:"token"`
}

func (config appConfig) Validate() error {
	return validation.ValidateStruct(&config,
		validation.Field(&config.ServerPort, validation.Required),
		validation.Field(&config.Hostname, validation.Required),
		validation.Field(&config.Database, validation.NotNil),
		validation.Field(&config.Port, validation.Required),
		validation.Field(&config.Password, validation.Required),
		validation.Field(&config.DbDsn, validation.Required),
		validation.Field(&config.Token, validation.Required),
	)
}

func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("app")
	v.SetConfigType("yaml")
	v.SetDefault("server_port", 8080)
	v.SetDefault("redis_hostname", "127.0.0.1")
	v.SetDefault("redis_port", 6379)
	v.SetDefault("redis_database", 0)
	v.SetDefault("redis_password", 0)
	v.SetDefault("db_data_dsn", 0)
	v.SetDefault("token", 0)
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
