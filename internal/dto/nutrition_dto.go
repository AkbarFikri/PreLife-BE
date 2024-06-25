package dto

type NutritionRequest struct {
	FoodName     string  `json:"food_name"`
	Calories     float64 `json:"calories"`
	Carbohidrate float64 `json:"carbohydrates"`
	Protein      float64 `json:"protein"`
}
