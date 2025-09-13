// Package models contains models for the Poll API
package models

import (
	"time"

	"github.com/google/uuid"
)

type PollStatus int

const (
	Created   PollStatus = iota // created
	Finalized                   // finalized
	Cancelled                   // cancelled
)

type Poll struct {
	ID        string     `json:"id"`
	PollName  string     `json:"pollName"`
	CreatorID string     `json:"creatorId"`
	CreatedAt string     `json:"createdAt"`
	ExpireSec int        `json:"expirationSeconds"`
	ExpireMin int        `json:"expirationMinutes"`
	Status    PollStatus `json:"status"`
}

func NewPoll(pollName string, creatorID string, expireSec int, expireMin int) *Poll {
	p := new(Poll)
	p.ID = uuid.NewString()
	p.PollName = pollName
	p.CreatorID = creatorID
	p.ExpireMin = expireMin
	p.ExpireSec = expireSec
	p.CreatedAt = time.Now().Format("01-02-2006 15:04:05")
	p.Status = Created
	return p
}

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"CreatedAt"`
}

func NewUser(userName string) *User {
	u := new(User)
	u.ID = uuid.NewString()
	u.Name = userName
	u.CreatedAt = time.Now().Format("01-02-2006 15:04:05")
	return u
}
