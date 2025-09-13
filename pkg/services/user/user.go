// Package user contains functions for the user type
package user

import (
	"context"
	"encoding/json"
	"fmt"

	mapper "example.com/go-polls/pkg/helpers"
	models "example.com/go-polls/pkg/models"
	redis "example.com/go-polls/pkg/services/redis"
)

func CreateUser(request models.CreateUserRequest) (string, error) {
	key := "user:" + request.Name

	newUser := models.NewUser(request.Name)

	params := mapper.UserToMap(newUser)

	err := redis.SetHash(context.Background(), key, params, 60*60)
	if err != nil {
		return "", fmt.Errorf("failed to save user in Redis: %w", err)
	}

	jsonBytes, err := json.Marshal(params)
	if err != nil {
		return "", fmt.Errorf("failed to marshal user data: %w", err)
	}

	return string(jsonBytes), nil
}

func GetUser(id string) (string, error) {
	val, err := redis.GetHash(context.Background(), id)
	if err != nil {
		return "", fmt.Errorf("failed to get user in Redis: %w", err)
	}

	jsonBytes, err := json.Marshal(val)
	if err != nil {
		return "", fmt.Errorf("failed to marshal user data: %w", err)
	}

	return string(jsonBytes), nil
}

func GetAllUsers() (string, error) {
	val, err := redis.GetAllKeys(context.Background(), "user:*")

	if err != nil {
		return "", fmt.Errorf("failed to get user in Redis: %w", err)
	}

	jsonBytes, err := json.Marshal(val)
	if err != nil {
		return "", fmt.Errorf("failed to marshal user data: %w", err)
	}

	return string(jsonBytes), nil
}
