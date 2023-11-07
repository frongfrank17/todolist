package configs

import (
	"fmt"
	"log"
	"os"

	env "github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	NODE_ENV           string `env:"NODE_ENV"`
	Name               string `env:"NAME"`
	PORT               string `env:"PORT"`
	Postgress_username string `env:"POSTGRESS_USERNAME"`
	Postgress_password string `env:"POSTGRESS_PASSWORD"`
	Postgress_host     string `env:"POSTGRESS_HOST"`
	Postgress_port     string `env:"POSTGRESS_PORT"`
	Postgress_db       string `env:"POSTGRESS_DB"`
}

func NewConfig() (*Config, error) {
	cfg := Config{}
	fmt.Println("NODE_ENV", os.Getenv("NODE_ENV"))
	if os.Getenv("NODE_ENV") == "Development" {

		err := godotenv.Load()

		if err != nil {
			log.Fatal("Error loading .env file")
			return nil, err
		}

		err = env.Parse(&cfg) // 👈 Parse environment variables into `Config`
		if err != nil {
			log.Fatalf("unable to parse ennvironment variables: %e", err)
			return nil, err
		}
		return &cfg, nil
	}

	err := env.Parse(&cfg) // 👈 Parse environment variables into `Config`
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
		return nil, err
	}
	return &cfg, nil

}
