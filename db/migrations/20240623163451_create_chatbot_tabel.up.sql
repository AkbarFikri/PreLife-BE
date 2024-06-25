CREATE TABLE IF NOT EXISTS chatbots (
    id SERIAL PRIMARY KEY,
    user_profile_id VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    response TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_user
        FOREIGN KEY(user_profile_id)
        REFERENCES user_profiles(id)
);