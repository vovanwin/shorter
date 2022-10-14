package config

type Config struct {
	IsProduction bool   `env:"PRODUCTION"`
	Domain       string `env:"HOSTS" envSeparator:":" envDefault:"127.0.0.1:8080"`
}
