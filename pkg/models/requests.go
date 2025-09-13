package models

type CreateUserRequest struct {
	Name string `json:"name"`
}

type CreatePollRequest struct {
	PollID    string `json:"pollID"`
	PollName  string `json:"pollName"`
	CreatorID string `json:"creatorID"`
	ExpireSec int    `json:"expirationSeconds"`
	ExpireMin int    `json:"expirationMinutes"`
}
