package services

import (
	"context"
	"fmt"
	models "example.com/go-polls/pkg/models"
	"github.com/redis/go-redis/v9"
)

func CreateUser(models.User)(bool, error)  {
	// TODO	
	return false, nil;
}

func GetUser(int)(string, error)  {
	// TODO	
	return "abc", nil;
}

func GetUsers()(string,error) {
	// TODO	
	return "abc", nil;
}
