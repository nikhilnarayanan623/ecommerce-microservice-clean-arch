package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServicePort       string `mapstructure:"CART_SERVICE_PORT"`
	ProductServiceUrl string `mapstructure:"PRODUCT_SERVICE_URL"`

	DBHost     string `mapstructure:"CART_DB_HOST"`
	DBPort     string `mapstructure:"CART_DB_PORT"`
	DBName     string `mapstructure:"CART_DB_NAME"`
	DBUser     string `mapstructure:"CART_DB_USER"`
	DBPassword string `mapstructure:"CART_DB_PASSWORD"`
}

var envs = []string{
	"CART_SERVICE_PORT", "PRODUCT_SERVICE_URL",
	"CART_DB_HOST", "CART_DB_PORT", "CART_DB_NAME", "CART_DB_USER", "CART_DB_PASSWORD",
}

func LoadConfig() (config *Config, err error) {

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")

	viper.ReadInConfig()

	for _, env := range envs {
		if err = viper.BindEnv(env); err != nil {
			return
		}
	}
	err = viper.Unmarshal(&config)

	return
}
