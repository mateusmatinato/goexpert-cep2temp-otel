package config

import (
	"github.com/mateusmatinato/goexpert-cep2temp-otel/internal/input/orchestration"
	"github.com/spf13/viper"
)

type Config struct {
	OrchestrationURL string `mapstructure:"orchestration_url"`
}

func (c *Config) OrchClientConfig() orchestration.APIConfig {
	return orchestration.APIConfig{OrchestrationURL: c.OrchestrationURL}
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
