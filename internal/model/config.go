package model

type Config struct {
	DbURL string `env:"DB_URL"`
	Port  string `env:"PORT"`
}
