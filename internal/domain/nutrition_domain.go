package domain

import "time"

type Nutrition struct {
	ID            string    `db:"id"`
	UserProfileId string    `db:"user_profile_id"`
	FoodName      string    `db:"food_name"`
	Calories      float64   `db:"calories"`
	Protein       float64   `db:"protein"`
	Carbohydrate  float64   `db:"carbohydrate"`
	PictureLink   string    `db:"picture_link"`
	CreatedAt     time.Time `db:"created_at"`
}
