package chatbotRepository

import (
	"github.com/AkbarFikri/PreLife-BE/internal/domain"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

type ChatbotRepository interface {
	Save(ctx context.Context, chat domain.Chatbot) (int, error)
	FindAllChatByProfileId(ctx context.Context, id string) ([]domain.Chatbot, error)
}

type chatbotRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) ChatbotRepository {
	return chatbotRepository{
		db: db,
	}
}

func (r chatbotRepository) Save(ctx context.Context, chat domain.Chatbot) (int, error) {
	arg := map[string]interface{}{
		"user_profile_id": chat.UserProfileId,
		"message":         chat.Message,
		"response":        chat.Response,
		"created_at":      chat.CreatedAt,
	}

	query, args, err := sqlx.Named(createChatbotsRecord, arg)
	if err != nil {
		return -1, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return -1, err
	}
	query = r.db.Rebind(query)

	var id int
	if err := r.db.QueryRowxContext(ctx, query, args...).Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func (r chatbotRepository) FindAllChatByProfileId(ctx context.Context, id string) ([]domain.Chatbot, error) {
	arg := map[string]interface{}{
		"user_profile_id": id,
	}

	query, args, err := sqlx.Named(getChatbotsRecordByProfileId, arg)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}
	query = r.db.Rebind(query)

	rows, err := r.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var chatHistory []domain.Chatbot
	for rows.Next() {
		var chat domain.Chatbot
		if err := rows.StructScan(&chat); err != nil {
			return chatHistory, err
		}
		chatHistory = append(chatHistory, chat)
	}

	return chatHistory, nil
}
