package dto

import "time"

type CheckIsPregnant struct {
	ProfileType float64 `form:"profile-type" binding:"required"`
}

type RegisterProfilePregnantRequest struct {
	UserID       string `json:"user_id" binding:"required"`
	IsPregnant   bool
	ProfileName  string `json:"profile_name" binding:"required"`
	PregnantDate string `json:"pregnant_date" binding:"required"`
}

type RegisterProfileNotPregnantRequest struct {
	UserID      string `json:"user_id" binding:"required"`
	IsPregnant  bool
	ProfileName string  `json:"profile_name" binding:"required"`
	BirthDate   string  `json:"birth_date" binding:"required"`
	Weight      float64 `json:"weight" binding:"required"`
	Height      float64 `json:"height" binding:"required"`
}

type RegisterProfileResponse struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	IsPregnant bool   `json:"is_pregnant"`
}

type PregnantProfileResponse struct {
	Id                 string    `json:"id"`
	UserID             string    `json:"user_id"`
	ProfileName        string    `json:"profile_name"`
	IsPregnant         bool      `json:"is_pregnant"`
	PregnantDate       time.Time `json:"pregnant_date"`
	PregnantAgeInDay   int       `json:"pregnant_age_in_day"`
	PregnantAgeInWeek  int       `json:"pregnant_age_in_week"`
	PregnantAgeInMonth int       `json:"pregnant_age_in_month"`
}

type NonPregnantProfileResponse struct {
	Id              string    `json:"id"`
	UserID          string    `json:"user_id"`
	ProfileName     string    `json:"profile_name"`
	IsPregnant      bool      `json:"is_pregnant"`
	BirthDate       time.Time `json:"birth_date"`
	Height          float64   `json:"height"`
	Weight          float64   `json:"weight"`
	ChildAgeInDay   int       `json:"child_age_in_day"`
	ChildAgeInWeek  int       `json:"child_age_in_week"`
	ChildAgeInMonth int       `json:"child_age_in_month"`
	ChildAgeInYear  int       `json:"child_age_in_year"`
}

type ProfileListResponse struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	IsPregnant  bool   `json:"is_pregnant"`
	ProfileName string `json:"profile_name"`
}
