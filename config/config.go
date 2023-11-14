package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	EmailConfig *EmailConfig
}

type EmailConfig struct {
	HOST string `envconfig:"EMAIL_HOST"`
	PORT int    `envconfig:"EMAIL_PORT" default:"25"`
	AUTH struct {
		USER     string `envconfig:"EMAIL_USERNAME"`
		PASSWORD string `envconfig:"EMAIL_PASSWORD"`
	}
}

func Configurations(envPath string) (Config, error) {
	cnf := Config{}
	err := godotenv.Load(envPath)
	if err != nil {
		return cnf, err
	}
	err = envconfig.Process("", &cnf)
	return cnf, err
}
