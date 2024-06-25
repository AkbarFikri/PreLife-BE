CREATE TABLE IF NOT EXISTS user_profiles (
    id VARCHAR(255) NOT NULL PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    is_pregnant BOOLEAN NOT NULL,
    pregnant_date TIMESTAMP,
    birth_date TIMESTAMP,
    profile_name VARCHAR(255) NOT NULL,
    weight INTEGER,
    height INTEGER,
    medical_center_id VARCHAR,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (medical_center_id) REFERENCES medical_centers(id)
);