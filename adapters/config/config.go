package config

import (
	"fmt"
	env "github.com/joho/godotenv"
	"log"
)

const (
	ENV_MODE string = "dev"
)

func init() {
	path := fmt.Sprintf("internals/config/%s.env", ENV_MODE)
	errLoadEnv := env.Load(path)

	if errLoadEnv != nil {
		log.Fatal("Error Loading .env file")
	}
}
