package domain

import "time"

type UserProfile struct {
	ID          string `db:"id"`
	UserID      string `db:"user_id"`
	ProfileName string `db:"profile_name"`
	IsPregnant  bool   `db:"is_pregnant"`
}

type PregnantProfile struct {
	ID           string    `db:"id"`
	UserId       string    `db:"user_id"`
	IsPregnant   bool      `db:"is_pregnant"`
	ProfileName  string    `db:"profile_name"`
	PregnantDate time.Time `db:"pregnant_date"`
}

type NotPregnantProfile struct {
	ID          string    `db:"id"`
	UserId      string    `db:"user_id"`
	IsPregnant  bool      `db:"is_pregnant"`
	BirthDate   time.Time `db:"birth_date"`
	ProfileName string    `db:"profile_name"`
	Weight      float64   `db:"weight"`
	Height      float64   `db:"height"`
}
