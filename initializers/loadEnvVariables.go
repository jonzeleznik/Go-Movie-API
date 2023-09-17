package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

var EnvVars map[string]string

func LoadEnvVariables() {
	var err error
	EnvVars, err = godotenv.Read()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
