package nutritionRepository

import (
	"github.com/AkbarFikri/PreLife-BE/internal/domain"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

type NutritionRepository interface {
	Save(ctx context.Context, data domain.Nutrition) error
	GetCurrentDateHistory(ctx context.Context, profileId string) ([]domain.Nutrition, error)
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

func (r nutritionRepository) GetCurrentDateHistory(ctx context.Context, profileId string) ([]domain.Nutrition, error) {
	arg := map[string]interface{}{
		"user_profile_id": profileId,
	}

	query, args, err := sqlx.Named(FindNutritionsInCurrentDateByProfileId, arg)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}
	query = r.db.Rebind(query)

	rows, err := r.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var nutritionsHistory []domain.Nutrition
	for rows.Next() {
		var nutrition domain.Nutrition
		if err := rows.StructScan(&nutrition); err != nil {
			return nutritionsHistory, err
		}
		nutritionsHistory = append(nutritionsHistory, nutrition)
	}

	return nutritionsHistory, nil
}
