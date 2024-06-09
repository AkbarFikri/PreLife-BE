package userRepository

const GetById = `
SELECT
	id,
	full_name,
	email
FROM
    users
WHERE
    id = :id
`
const CreateUser = `
INSERT INTO
	users (
	       id,
	       email,
	       full_name
) VALUES (
          :id,
          :email,
          :full_name
)`

const CountEmail = `
SELECT 
	COUNT(email) 
FROM 
	users 
WHERE 
	email = :email
`
