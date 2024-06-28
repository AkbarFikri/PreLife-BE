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

const FindNutritionsInCurrentDateByProfileId = `
SELECT
	id,
	user_profile_id,
	calories,
	carbohydrate,
	protein
FROM
	nutritions
WHERE
    DATE(created_at) = CURRENT_DATE AND user_profile_id = :user_profile_id`
