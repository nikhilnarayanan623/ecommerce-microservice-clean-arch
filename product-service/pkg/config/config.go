package config

import "github.com/spf13/viper"

type Config struct {
	ServicePort string `mapstructure:"PRODUCT_SERVICE_PORT"`

	DBHost     string `mapstructure:"PRODUCT_DB_HOST"`
	DBPort     string `mapstructure:"PRODUCT_DB_PORT"`
	DBName     string `mapstructure:"PRODUCT_DB_NAME"`
	DBUser     string `mapstructure:"PRODUCT_DB_USER"`
	DBPassword string `mapstructure:"PRODUCT_DB_PASSWORD"`
}

var envs = []string{
	"PRODUCT_SERVICE_PORT",
	"PRODUCT_DB_HOST", "PRODUCT_DB_PORT", "PRODUCT_DB_NAME", "PRODUCT_DB_USER", "PRODUCT_DB_PASSWORD",
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
