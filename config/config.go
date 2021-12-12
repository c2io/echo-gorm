package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Configuration struct {
	Database struct {
		User     string `env:"DB_USER"`
		Password string `env:"DB_PASSWORD"`
		Port     int    `env:"DB_PORT"`
		Host     string `env:"DB_HOST"`
		Name     string `toml:"name" env:"DB_NAME"`
	}
}

func GetConfig() Configuration {
	c := Configuration{}
	cleanenv.ReadConfig("config/config.toml", &c)

	return c
}
