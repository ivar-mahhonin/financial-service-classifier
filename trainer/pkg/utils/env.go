package util

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvFile() error {
	dir, _ := os.Getwd()
	err := godotenv.Load(fmt.Sprintf("%s/.env", dir))
	return err
}

func GetEnvVariable(key string) string {
	return os.Getenv(key)
}
