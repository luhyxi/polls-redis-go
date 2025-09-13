package helpers

import (
	"strconv"

	models "example.com/go-polls/pkg/models"
)

func UserToMap(u *models.User) map[string]string {
	return map[string]string{
		"id":        u.ID,
		"name":      u.Name,
		"createdAt": u.CreatedAt,
	}
}

func PollToMap(p *models.Poll) map[string]string {
	return map[string]string{
		"id":        p.ID,
		"pollName":  p.PollName,
		"creatorId": p.CreatorID,
		"createdAt": p.CreatedAt,
		"status":    strconv.Itoa(int(p.Status)),
	}
}
