package util

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVariable(key string) string {
	dir, _ := os.Getwd()
	err := godotenv.Load(fmt.Sprintf("%s/.env", dir))
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
