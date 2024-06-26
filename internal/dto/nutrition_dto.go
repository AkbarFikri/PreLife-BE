package dto

import "mime/multipart"

type GenerateNutritionsResponse struct {
	FoodName     string  `json:"food_name"`
	Calories     float64 `json:"calories"`
	Carbohidrate float64 `json:"carbohydrates"`
	Protein      float64 `json:"protein"`
	Recommended  bool    `json:"recommended_for_pregnant_children"`
	Description  string  `json:"description"`
}

type GenerateNutritionRequest struct {
	Picture *multipart.FileHeader `form:"picture" binding:"required"`
}

type NutritionsRequest struct {
	Picture      *multipart.FileHeader `form:"picture" binding:"required"`
	FoodName     string                `form:"food_name" binding:"required"`
	Calories     float64               `form:"calories" binding:"required"`
	Carbohidrate float64               `form:"carbohydrates" binding:"required"`
	Protein      float64               `form:"protein" binding:"required"`
}

type NutritionsResponse struct {
	ID          string `json:"id"`
	FoodName    string `json:"food_name"`
	PictureLink string `json:"picture_link"`
}
