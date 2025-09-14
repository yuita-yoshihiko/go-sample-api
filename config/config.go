package config

type Config struct {
	DBConfig
}

var Conf Config

type DBConfig struct {
	DatabaseURL string `env:"DATABASE_URL,required"`
}
