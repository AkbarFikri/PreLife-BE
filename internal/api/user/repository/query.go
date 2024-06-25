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

const GetUserById = `
SELECT
	id,
	full_name,
	email,
	date_of_birth,
	role_id
FROM
	users
WHERE
    id = :id`

const GetPregnantProfileById = `
SELECT
	id,
	profile_name,
	is_pregnant,
	user_id,
	pregnant_date
FROM
	user_profiles
WHERE
	id = :id`

const GetNonPregnantProfileById = `
SELECT
	id,
	profile_name,
	is_pregnant,
	user_id,
	birth_date,
	height,
	weight
FROM
	user_profiles
WHERE
	id = :id`

const GetAllProfilesByUserId = `
SELECT
	id,
	user_id,
	profile_name,
	is_pregnant
FROM
	user_profiles
WHERE
    user_id = :user_id`

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

const CreateProfilPregnant = `
INSERT INTO 
    user_profiles (
    	id,
        user_id,
    	profile_name,
        is_pregnant,
        pregnant_date
) VALUES (
    	:id,
        :user_id,
        :profile_name,
        :is_pregnant,
        :pregnant_date
)`

const CreateProfilNotPregnant = `
INSERT INTO 
    user_profiles (
    	id,
        user_id,
    	profile_name,
        is_pregnant,
        birth_date,
        height,
    	weight
) VALUES (
    	:id,
        :user_id,
        :profile_name,
        :is_pregnant,
        :birth_date,
        :height,
        :weight
)`
