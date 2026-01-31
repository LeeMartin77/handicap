package config

type Config struct {
	Port        string `env:"PORT"`
	PostgresUrl string `env:"POSTGRES_URL"`
}
