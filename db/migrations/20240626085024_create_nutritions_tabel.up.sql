CREATE TABLE IF NOT EXISTS nutritions (
    id VARCHAR NOT NULL PRIMARY KEY,
    user_profile_id VARCHAR NOT NULL,
    food_name VARCHAR NOT NULL,
    calories INTEGER NOT NULL,
    carbohydrate INTEGER NOT NULL,
    protein INTEGER NOT NULL,
    picture_link TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_profile_id) REFERENCES user_profiles(id)
);