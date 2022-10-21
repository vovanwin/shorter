package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	ServerAddress       string `env:"SERVER_ADDRESS,required" envSeparator:":" envDefault:"127.0.0.1:8080"`
	BaseURL             string `env:"BASE_URL,required" envDefault:"/api/shorten"`
	FileStoragePath     string `env:"FILE_STORAGE_PATH" envDefault:"./storage/url.json"`
	FileStoragePathTest string `env:"FILE_STORAGE_PATH" envDefault:"./../../../storage/test.json"`
}

func (c Config) GetConfig() Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
	return cfg
}
