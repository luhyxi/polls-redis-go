package models

type CreateUserRequest struct {
	Name      string `json:"name"`
}

type CreatePollRequest struct {
	Name      string `json:"name"`
	ExpireSec int    `json:"expirationSeconds"`
	ExpireMin int    `json:"expirationMinutes"`
}

