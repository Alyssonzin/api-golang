package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Envs struct {
	AppPort            string
	DBHost             string
	DBPort             string
	DBUser             string
	DBPassword         string
	DBName             string
	PluggyClientID     string
	PluggyClientSecret string
}

func LoadEnvs(path string) Envs {
	if err := godotenv.Load(path); err != nil {
		panic(err)
	}

	return Envs{
		AppPort:            os.Getenv("APP_PORT"),
		DBHost:             os.Getenv("DB_HOST"),
		DBPort:             os.Getenv("DB_PORT"),
		DBUser:             os.Getenv("DB_USER"),
		DBPassword:         os.Getenv("DB_PASSWORD"),
		DBName:             os.Getenv("DB_NAME"),
		PluggyClientID:     os.Getenv("PLUGGY_CLIENT_ID"),
		PluggyClientSecret: os.Getenv("PLUGGY_CLIENT_SECRET"),
	}
}
