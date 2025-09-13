// Package poll contains functions for the poll type
package poll

import (
	"context"
	"encoding/json"
	"fmt"

	mapper "example.com/go-polls/pkg/helpers"
	"example.com/go-polls/pkg/models"
	redis "example.com/go-polls/pkg/services/redis"
)

func CreatePoll(request models.CreatePollRequest) (string, error) {
	key := "poll:" + request.PollName

	newPoll := models.NewPoll(
		request.PollName,
		request.CreatorID,
		request.ExpireSec,
		request.ExpireMin,
	)

	params := mapper.PollToMap(newPoll)

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

func GetPoll(id string) (string, error) {
	//TODO
	return "", nil
}

func GetAllPolls() (string, error) {
	//TODO
	return "", nil
}
