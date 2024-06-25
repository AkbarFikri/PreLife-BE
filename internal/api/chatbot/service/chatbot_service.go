package chatbotService

import (
	chatbotRepository "github.com/AkbarFikri/PreLife-BE/internal/api/chatbot/repository"
	"github.com/AkbarFikri/PreLife-BE/internal/domain"
	"github.com/AkbarFikri/PreLife-BE/internal/dto"
	"github.com/AkbarFikri/PreLife-BE/internal/pkg/gemini"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"time"
)

type ChatbotService interface {
	GenerateChatbotResponseFrom(ctx context.Context, user dto.UserTokenData, message string) (dto.ChatResponse, error)
	GetProfileChatbotHistory(ctx context.Context, user dto.UserTokenData) ([]dto.ChatResponse, error)
}

type chatbotService struct {
	log               *logrus.Logger
	gemini            *gemini.Gemini
	chatbotRepository chatbotRepository.ChatbotRepository
}

func New(l *logrus.Logger, gemini *gemini.Gemini, cs chatbotRepository.ChatbotRepository) ChatbotService {
	return &chatbotService{
		log:               l,
		gemini:            gemini,
		chatbotRepository: cs,
	}
}

func (s chatbotService) GenerateChatbotResponseFrom(ctx context.Context, user dto.UserTokenData, message string) (dto.ChatResponse, error) {
	res, err := s.gemini.GenerateChatResponse(ctx, message)
	if err != nil {
		return dto.ChatResponse{}, err
	}

	chatbot := domain.Chatbot{
		UserProfileId: user.ProfileId,
		Message:       message,
		Response:      res,
		CreatedAt:     time.Now(),
	}

	id, err := s.chatbotRepository.Save(ctx, chatbot)
	if err != nil {
		s.log.Errorf("error when saving chatbot: %v", err)
		return dto.ChatResponse{}, err
	}

	return dto.ChatResponse{
		ID:        id,
		Response:  res,
		CreatedAt: time.Now(),
		Message:   message,
	}, nil
}

func (s chatbotService) GetProfileChatbotHistory(ctx context.Context, user dto.UserTokenData) ([]dto.ChatResponse, error) {
	history, err := s.chatbotRepository.FindAllChatByProfileId(ctx, user.ProfileId)
	if err != nil {
		return nil, err
	}

	if len(history) == 0 {
		return []dto.ChatResponse{}, nil
	}

	var res []dto.ChatResponse

	for _, h := range history {
		res = append(res, dto.ChatResponse{
			ID:        h.ID,
			Message:   h.Message,
			CreatedAt: h.CreatedAt,
			Response:  h.Response,
		})
	}
	return res, nil
}
