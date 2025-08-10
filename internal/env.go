package internal

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetRedisURL() (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not get redis url")
		return "", errors.New("could not get redis url")
	}

	return os.Getenv("REDIS_URL"), nil
}
