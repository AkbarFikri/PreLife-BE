package nutritionRepository

import (
	"github.com/AkbarFikri/PreLife-BE/internal/domain"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

type NutritionRepository interface {
	Save(ctx context.Context, data domain.Nutrition) error
}

type nutritionRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) NutritionRepository {
	return nutritionRepository{db: db}
}

func (r nutritionRepository) Save(ctx context.Context, data domain.Nutrition) error {
	arg := map[string]interface{}{
		"id":              data.ID,
		"user_profile_id": data.UserProfileId,
		"food_name":       data.FoodName,
		"calories":        data.Calories,
		"protein":         data.Protein,
		"carbohydrate":    data.Carbohydrate,
		"picture_link":    data.PictureLink,
		"created_at":      data.CreatedAt,
	}

	_, err := r.db.NamedExecContext(ctx, CreateNutritions, arg)
	if err != nil {
		return err
	}
	return nil
}
