package config

import "github.com/spf13/viper"

type Config struct {
	ServicePort string `mapstructure:"USER_SERVICE_PORT"`
	DBHost      string `mapstructure:"USER_DB_HOST"`
	DBPort      string `mapstructure:"USER_DB_PORT"`
	DBName      string `mapstructure:"USER_DB_NAME"`
	DBUser      string `mapstructure:"USER_DB_USER"`
	DBPassword  string `mapstructure:"USER_DB_PASSWORD"`
}

var envs = []string{
	"USER_SERVICE_PORT",
	"USER_DB_HOST", "USER_DB_PORT", "USER_DB_NAME", "USER_DB_USER", "USER_DB_PASSWORD",
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
