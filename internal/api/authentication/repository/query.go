package authRepository

const CreateUser = `
INSERT INTO
	users (
	       id,
	       email,
	       full_name,
	       date_of_birth,
	       role_id
) VALUES (
          :id,
          :email,
          :full_name,
          :date_of_birth,
          :role_id
)`

const GetUserByEmail = `
SELECT
	id,
	full_name,
	email,
	role_id,
	roles.role_name
FROM
	users
JOIN roles ON roles.role_id = users.role_id
WHERE
    email = :email`

const CountEmail = `
SELECT 
	COUNT(email) 
FROM 
	users 
WHERE 
	email = :email
`

const GetUserProfileIdByUserId = `
SELECT
	id
FROM
    user_profiles
WHERE
    user_id = :user_id
`
