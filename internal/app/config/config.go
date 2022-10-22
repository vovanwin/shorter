package config

import (
	"flag"
	"fmt"
	"github.com/caarlos0/env/v6"
)

var (
	ServerAddress   *string
	BaseURL         *string
	FileStoragePath *string
)

func init() {
	ServerAddress = flag.String("a", "127.0.0.1:8080", "SERVER_ADDRESS")
	BaseURL = flag.String("b", "/api/shorten", "BASE_URL")
	FileStoragePath = flag.String("f", "./storage/url.json", "FILE_STORAGE_PATH")
}

type Config struct {
	ServerAddress       string `env:"SERVER_ADDRESS,required" envSeparator:":" envDefault:"127.0.0.1:8080"`
	BaseURL             string `env:"BASE_URL,required" envDefault:"/api/shorten"`
	FileStoragePath     string `env:"FILE_STORAGE_PATH" envDefault:"./storage/url.json"`
	FileStoragePathTest string `env:"FILE_STORAGE_PATH" envDefault:"./../../../storage/test.json"`
}

// GetConfig TODO: возможно надо сделать синглтон
func (c Config) GetConfig() Config {
	if err := env.Parse(&c); err != nil {
		fmt.Printf("%+v\n", err)
	}

	c.ServerAddress = *ServerAddress
	c.BaseURL = *BaseURL
	c.FileStoragePath = *FileStoragePath
	return c
}
