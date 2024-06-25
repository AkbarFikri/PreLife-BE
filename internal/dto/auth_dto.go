package dto

import "time"

type AuthRequest struct {
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=8"`
	FullName    string `json:"full_name" binding:"required,max=255"`
	DateOfBirth string `json:"date_of_birth" binding:"required,max=255"`
}

type AuthResponse struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	FullName    string `json:"full_name"`
	DateOfBirth string `json:"date_of_birth"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

type UserTokenData struct {
	ProfileId   string  `json:"profile_id"`
	ID          string  `json:"user_id"`
	Email       string  `json:"email"`
	RoleId      float64 `json:"role_id"`
	ProfileType float64 `json:"profile_type"`
}
