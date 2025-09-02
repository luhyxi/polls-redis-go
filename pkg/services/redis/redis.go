// Package redis contains the redis service
package redis

import (
	"context"
	"log"
	"time"

	"example.com/go-polls/internal"
	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func init() {
	url, err := internal.GetRedisURL()
	if err != nil {
		log.Fatal("Failed to get Redis URL:", err)
	}

	opts, err := redis.ParseURL(url)
	if err != nil {
		log.Fatal("Failed to parse Redis URL:", err)
	}

	rdb = redis.NewClient(opts)

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
}

func GetClient() *redis.Client {
	return rdb
}

func SetKeyValue(key string, value string, seconds time.Duration) bool {
	set, err := rdb.SetNX(ctx, key, value, seconds*time.Second).Result()
	if err != nil {
		log.Print("Unable to set key", err)
	}
	return set
}
