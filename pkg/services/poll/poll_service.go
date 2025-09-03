package poll

import (
	"context"
	"fmt"

	models "example.com/go-polls/pkg/models"
	redis "example.com/go-polls/pkg/services/redis"
	"github.com/google/uuid"
)

func CreateUser(user models.User) (bool, error) {
	userId := uuid.NewString()

    key := "user:" + userId

    params := map[string]string{
        "id":        userId,
        "name":      user.Name,
        "createdAt": user.CreatedAt,
    }

	err := redis.SetHash(context.Background(), key, params, 60*60)
    if err != nil {
        return false, fmt.Errorf("failed to save user in Redis")
    }

    return true, nil
}


func GetUser(int)(string, error)  {
	// TODO	
	return "abc", nil;
}

func GetUsers()(string,error) {
	// TODO	
	return "abc", nil;
}
