package domain

import "time"

type Chatbot struct {
	ID            int       `db:"id"`
	UserProfileId string    `db:"user_profile_id"`
	Message       string    `db:"message"`
	Response      string    `db:"response"`
	CreatedAt     time.Time `db:"created_at"`
}
