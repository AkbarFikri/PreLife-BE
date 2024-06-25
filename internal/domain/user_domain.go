package domain

import "time"

type User struct {
	ID          string
	Email       string
	FullName    string    `db:"full_name"`
	DateOfBirth time.Time `db:"date_of_birth"`
	RoleId      int       `db:"role_id"`
	RoleName    string    `db:"role_name"`
}
