package util

import "github.com/spf13/viper"

// Config stores all configuration of the application.
// The values are ready by viper from a config file or environment variables
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER" envconfig:"DB_DRIVER" default:"postgres"`
	DBSource      string `mapstructure:"DB_SOURCE" envconfig:"DB_SOURCE" default`
	ServerAddress string `mapstructure:"SERVER_ADDRESS" envconfig:"SERVER_ADDRESS" default:":8080"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}