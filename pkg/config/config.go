package config

import (
	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

type Config struct {
	Port       string `mapstructure:"PORT"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBPassword string `mapstructure:"DB_PASSWORD"`

	AwsAccessKeyID string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AwsSecretKey   string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	AwsRegion      string `mapstructure:"AWS_REGION"`
	AwsBucketName  string `mapstructure:"AWS_BUCKET_NAME"`
}

var envs = []string{
	"PORT", "DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD",
	"AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY", "AWS_REGION", "AWS_BUCKET_NAME",
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	return config, nil
}
