package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	ServerAddress   string `env:"SERVER_ADDRESS" envSeparator:":" envDefault:"127.0.0.1:8080"`
	BaseURL         string `env:"BASE_URL" envDefault:"/api/shorten"`
	FileStoragePath string `env:"FILE_STORAGE_PATH" envDefault:"./storage/"`
}

func NewConfig(serverAddress string, baseURL string, fileStoragePath string) *Config {
	return &Config{ServerAddress: serverAddress, BaseURL: baseURL, FileStoragePath: fileStoragePath}
}

func (c Config) GetConfig() Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
	return cfg
}
