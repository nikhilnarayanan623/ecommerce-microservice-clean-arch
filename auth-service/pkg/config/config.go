package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServicePort    string `mapstructure:"AUTH_SERVICE_PORT"`
	UserServiceUrl string `mapstructure:"USER_SERVICE_URL"`

	DBHost     string `mapstructure:"AUTH_DB_HOST"`
	DBPort     string `mapstructure:"AUTH_DB_PORT"`
	DBName     string `mapstructure:"AUTH_DB_NAME"`
	DBUser     string `mapstructure:"AUTH_DB_USER"`
	DBPassword string `mapstructure:"AUTH_DB_PASSWORD"`

	AdminSecretKey   string `mapstructure:"ADMIN_SECRET_KEY"`
	UserSecretKey    string `mapstructure:"USER_SECRET_KEY"`
	TwilioServiceID  string `mapstructure:"TWILIO_SERVICE_ID"`
	TwilioAuthToken  string `mapstructure:"TWILIO_AUTH_TOKEN"`
	TwilioAccountSID string `mapstructure:"TWILIO_ACCOUNT_SID"`
}

var envs = []string{
	"AUTH_SERVICE_PORT", "USER_SERVICE_URL",
	"AUTH_DB_HOST", "AUTH_DB_PORT", "AUTH_DB_NAME", "AUTH_DB_USER", "AUTH_DB_PASSWORD",
	"ADMIN_SECRET_KEY", "USER_SECRET_KEY",
	"TWILIO_SERVICE_ID", "TWILIO_AUTH_TOKEN", "TWILIO_ACCOUNT_SID",
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
