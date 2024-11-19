package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

// loads environment variables from the .env file
func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

//retrieves the value of an environment variable
func GetEnv(key string) string {
	return os.Getenv(key)
}
