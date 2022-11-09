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
	DatabaseDsn     *string
)

func init() {
	ServerAddress = flag.String("a", "", "SERVER_ADDRESS")
	BaseURL = flag.String("b", "", "BASE_URL")
	FileStoragePath = flag.String("f", "", "FILE_STORAGE_PATH")
	DatabaseDsn = flag.String("d", "", "DATABASE_DSN")
}

type Config struct {
	ServerAddress       string `env:"SERVER_ADDRESS,required" envSeparator:":" envDefault:"127.0.0.1:8080"`
	BaseURL             string `env:"BASE_URL,required" envDefault:"/api/shorten"`
	FileStoragePath     string `env:"FILE_STORAGE_PATH" envDefault:"./storage/url.json"`
	FileStoragePathTest string `env:"FILE_STORAGE_PATH" envDefault:"./../../../storage/test.json"`
	Key                 string `env:"KEY" envDefault:"hfjvjdjcdsdsdfdshfjvjdjcdsdsdfds"`
	DatabaseDsn         string `env:"DATABASE_DSN" envDefault:"postgres://default:secret@localhost:5432/shorter"`
}

// GetConfig TODO: возможно надо сделать синглтон
func (c Config) GetConfig() Config {
	if err := env.Parse(&c); err != nil {
		fmt.Printf("%+v\n", err)
	}
	if *ServerAddress != "" {
		c.ServerAddress = *ServerAddress
	}
	if *BaseURL != "" {
		c.BaseURL = *BaseURL
	}
	if *FileStoragePath != "" {
		c.FileStoragePath = *FileStoragePath
	}
	if *DatabaseDsn != "" {
		c.DatabaseDsn = *DatabaseDsn
	}

	return c
}
