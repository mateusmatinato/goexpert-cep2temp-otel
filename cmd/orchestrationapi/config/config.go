package config

import (
	"github.com/mateusmatinato/goexpert-cep2temp-otel/internal/orchestration/cep"
	"github.com/mateusmatinato/goexpert-cep2temp-otel/internal/orchestration/weather"
	"github.com/spf13/viper"
)

type Config struct {
	WeatherURL string `mapstructure:"weather_url"`
	WeatherKey string `mapstructure:"weather_api_key"`
	CepURL     string `mapstructure:"cep_url"`
}

func (c *Config) WeatherAPIConfig() weather.APIConfig {
	return weather.APIConfig{
		URL:    c.WeatherURL,
		APIKey: c.WeatherKey,
	}
}

func (c *Config) CepAPIConfig() cep.APIConfig {
	return cep.APIConfig{
		URL: c.CepURL,
	}
}

func LoadConfig(path string) (cfg Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&cfg)
	return
}
