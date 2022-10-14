package config

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS" envSeparator:":" envDefault:"127.0.0.1:8080"`
	BaseUrl       string `env:"BASE_URL" envDefault:"api/shorter"`
}
