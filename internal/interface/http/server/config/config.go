package config

type Config struct {
	Adr   string `env:"ADR_PATH"`
	Token string `env:"API_TOKEN"`
}
