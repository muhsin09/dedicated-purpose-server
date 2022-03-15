package objects

import "github.com/google/uuid"

type AddMessageRequest struct {
	Message string `json:"message"`
}

type PollMessageRequest struct {
	ID uuid.UUID `json:"ID"`
}
