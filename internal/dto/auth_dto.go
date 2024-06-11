package dto

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
