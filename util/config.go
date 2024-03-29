package util

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are ready by viper from a config file or environment variables
type Config struct {
	DBDriver             string        `mapstructure:"DB_DRIVER" envconfig:"DB_DRIVER" default:"postgres"`
	DBSource             string        `mapstructure:"DB_SOURCE" envconfig:"DB_SOURCE" default:""`
	HTTPServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS" envconfig:"HTTP_SERVER_ADDRESS" default:":8080"`
	GRPCServerAddress    string        `mapstructure:"GRPC_SERVER_ADDRESS" envconfig:"GRPC_SERVER_ADDRESS" default:":9090"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY" envconfig:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION" envconfig:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION" envconfig:"REFRESH_TOKEN_DURATION"`
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
