package helper

import (
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
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

func HashPasswords(plain string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(hashed)
}

func ComparePassword(hashed string, normal string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(normal))
}
