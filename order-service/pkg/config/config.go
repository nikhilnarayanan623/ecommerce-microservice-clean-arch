package config

import "github.com/spf13/viper"

type Config struct {
	ServicePort       string `mapstructure:"ORDER_SERVICE_PORT"`
	CartServiceUrl    string `mapstructure:"CART_SERVICE_URL"`
	ProductServiceUrl string `mapstructure:"PRODUCT_SERVICE_URL"`

	DBHost     string `mapstructure:"ORDER_DB_HOST"`
	DBPort     string `mapstructure:"ORDER_DB_PORT"`
	DBName     string `mapstructure:"ORDER_DB_NAME"`
	DBUser     string `mapstructure:"ORDER_DB_USER"`
	DBPassword string `mapstructure:"ORDER_DB_PASSWORD"`
}

var envs = []string{
	"ORDER_SERVICE_PORT", "CART_SERVICE_URL", "PRODUCT_SERVICE_URL",
	"ORDER_DB_HOST", "ORDER_DB_PORT", "ORDER_DB_NAME", "ORDER_DB_USER", "ORDER_DB_PASSWORD",
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
