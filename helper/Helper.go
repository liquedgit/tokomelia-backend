package helper

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GoDotEnvVariables(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("File not found")
	}
	return os.Getenv(key)
}
