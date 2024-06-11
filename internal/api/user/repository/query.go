package userRepository

const GetUserByEmail = `
SELECT
	id,
	full_name,
	email,
	date_of_birth
FROM
	users
WHERE
    email = :email`

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

const CountEmail = `
SELECT 
	COUNT(email) 
FROM 
	users 
WHERE 
	email = :email
`
