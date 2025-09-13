// Package redis contains the redis service
package redis

import (
	"context"
	"log"
	"time"

	"example.com/go-polls/internal"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

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
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
}

func SetKeyValue(ctx context.Context, key string, value string, seconds int) error {
	_, err := rdb.SetNX(ctx, key, value, time.Duration(seconds)*time.Second).Result()
	if err != nil {
		log.Print("unable to set key:", err)
		return err
	}
	return nil
}

func SetHash(ctx context.Context, key string, params map[string]string, seconds int) error {
	if err := rdb.HSet(ctx, key, params).Err(); err != nil {
		log.Print("unable to set hash:", err)
		return err
	}

	if seconds > 0 {
		if err := rdb.Expire(ctx, key, time.Duration(seconds)*time.Second).Err(); err != nil {
			log.Print("unable to set expiration:", err)
			return err
		}
	}
	return nil
}

func GetHash(ctx context.Context, key string) (map[string]string, error) {
	hash, err := rdb.HGetAll(ctx, key).Result()
	if err != nil {
		log.Print("unable to get hash:", err)
		return nil, err
	}
	return hash, nil
}

func GetAllKeys(ctx context.Context, pattrn string) ([]string, error) {
	hash, err := rdb.Keys(ctx, pattrn).Result()
	if err != nil {
		log.Print("unable to get keys:", err)
		return nil, err
	}

	return hash, nil
}
