package configs

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServiceAPort  string `mapstructure:"SERVICE_A_PORT"`
	ServiceBURL   string `mapstructure:"SERVICE_B_URL"`
	ServiceBPort  string `mapstructure:"SERVICE_B_PORT"`
	ZipkinURL     string `mapstructure:"ZIPKIN_URL"`
	ZipkinPort    string `mapstructure:"ZIPKIN_PORT"`
	WeatherApiKey string `mapstructure:"WEATHER_API_KEY"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg Config

	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.BindEnv("SERVICE_A_PORT")
	viper.BindEnv("SERVICE_B_URL")
	viper.BindEnv("SERVICE_B_PORT")
	viper.BindEnv("ZIPKIN_URL")
	viper.BindEnv("ZIPKIN_PORT")
	viper.BindEnv("WEATHER_API_KEY")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("WARNING: %v\n", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
