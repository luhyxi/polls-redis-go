package internal

import (
	"github.com/redis/go-redis/v9"
)

func healthcheck() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,  // Connection protocol
	})
}
