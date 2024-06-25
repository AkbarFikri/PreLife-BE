package chatbotRepository

const createChatbotsRecord = `
INSERT INTO chatbots (
    user_profile_id,
    message,
    response,
    created_at
) VALUES (
    :user_profile_id,
    :message,
    :response,
    :created_at
) RETURNING id
`

const getChatbotsRecordByProfileId = `
SELECT
	id,
	user_profile_id,
	response,
	message,
	created_at
FROM chatbots
WHERE user_profile_id = :user_profile_id
ORDER BY id ASC`
