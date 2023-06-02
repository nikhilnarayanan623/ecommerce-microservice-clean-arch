package config

import "github.com/spf13/viper"

type Config struct {
	ServiceUrl     string `mapstructure:"SERVICE_URL"`
	UserServiceUrl string `mapstructure:"USER_SERVICE_URL"`

	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`

	AdminSecretKey string `mapstructure:"ADMIN_SECRET_KEY"`
	UserSecretKey  string `mapstructure:"USER_SECRET_KEY"`

	TwilioServiceID  string `mapstructure:"TWILIO_SERVICE_ID"`
	TwilioAuthToken  string `mapstructure:"TWILIO_AUTH_TOKEN"`
	TwilioAccountSID string `mapstructure:"TWILIO_ACCOUNT_SID"`
}

var envs = []string{
	"SERVICE_URL", "USER_SERVICE_URL",
	"DB_HOST", "DB_PORT", "DB_NAME", "DB_USER", "DB_PASSWORD",
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
