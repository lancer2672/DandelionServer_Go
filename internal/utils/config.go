package utils

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver          string `mapstructure:"DB_DRIVER"`
	Environment       string `mapstructure:"ENVIRONMENT"`
	DBSource          string `mapstructure:"DB_SOURCE"`
	ServerAddress     string `mapstructure:"SERVER_ADDRESS"`
	GRPCServerAddress string `mapstructure:"GRPC_SERVER_ADDRESS"`
	AwsAccessKey      string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AwsSecretKey      string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	AwsAkid           string `mapstructure:"SERVER_ADDRESS"`
	MigrationUrl      string `mapstructure:"MIGRATION_URL"`
	GatewayApiKey     string `mapstructure:"GATEWAY_API_KEY"`
}

// overrided by env if exists
func LoadConfig(path string) (config Config, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)

	configProject(config.Environment)
	return
}

func configProject(env string) {
	if env == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
}
