package authRepository

const CreateUser = `
INSERT INTO
	users (
	       id,
	       email,
	       full_name,
	       date_of_birth
) VALUES (
          :id,
          :email,
          :full_name,
          :date_of_birth
)`

const GetUserByEmail = `
SELECT
	id,
	full_name,
	email
FROM
	users
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
