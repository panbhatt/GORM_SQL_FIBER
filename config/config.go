package config

import (
	"os"

	"github.com/rs/zerolog/log"

	"github.com/joho/godotenv"
)

// Config func to get env value
func Config(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Error().Msg("Error loading .env file")
	}
	return os.Getenv(key)
}
