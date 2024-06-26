package nutritionRepository

const CreateNutritions = `
INSERT INTO
	nutritions (
	       id,
	       user_profile_id,
	       food_name,
	       calories,
	       protein,
	       carbohydrate,
	       picture_link,
	       created_at
) VALUES (
	       :id,
	       :user_profile_id,
	       :food_name,
	       :calories,
	       :protein,
	       :carbohydrate,
	       :picture_link,
	       :created_at
)`
