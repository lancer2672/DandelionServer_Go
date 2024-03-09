package utils

import "github.com/spf13/viper"

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	AwsAccessKey string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AwsSecretKey string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	AwsAkid string `mapstructure:"SERVER_ADDRESS"`
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
	return
}
