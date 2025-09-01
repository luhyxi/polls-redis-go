// Package models contains models for the Poll API
package models

type PollStatus int

const (
	Created   PollStatus = iota // created
	Finalized                   // finalized
	Cancelled                   // cancelled
)

type Poll struct {
	ID        string     `json:"id"`
	CreatorID string     `json:"creatorId"`
	CreatedAt string     `json:"createdAt"`
	Status    PollStatus `json:"status"`
}

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"CreatedAt"`
}
