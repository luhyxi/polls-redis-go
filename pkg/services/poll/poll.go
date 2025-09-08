package poll

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	//models "example.com/go-polls/pkg/models"
	redis "example.com/go-polls/pkg/services/redis"
	"github.com/google/uuid"
)

func CreateUser(userName string) (bool, error) {
	userID := uuid.NewString()

    key := "user:" + userID
	now := time.Now()

    params := map[string]string{
        "id":        userID,
        "name":      userName,
        "createdAt": now.Format("01-02-2006 15:04:05"),
    }

	err := redis.SetHash(context.Background(), key, params, 60*60)
    if err != nil {
		return false, fmt.Errorf("failed to save user in Redis: %w", err)
    }

    return true, nil
}


func GetUser(id string)(string, error)  {
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

func GetUsers()(string,error) {
	// TODO	
	return "abc", nil;
}
