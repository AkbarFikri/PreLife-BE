package dto

import "time"

type ChatRequest struct {
	Message string `json:"message" binding:"required"`
}

type ChatResponse struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	Response  string    `json:"response"`
	CreatedAt time.Time `json:"created_at"`
}
