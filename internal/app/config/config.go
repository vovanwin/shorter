package config

type Config struct {
	SERVER_ADDRESS string `env:"SERVER_ADDRESS" envSeparator:":" envDefault:"127.0.0.1:8080"`
	BASE_URL       string `env:"BASE_URL" envDefault:"api/shorter"`
}
