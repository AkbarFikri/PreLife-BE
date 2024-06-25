package chatbotHandler

import (
	chatbotService "github.com/AkbarFikri/PreLife-BE/internal/api/chatbot/service"
	"github.com/AkbarFikri/PreLife-BE/internal/dto"
	"github.com/AkbarFikri/PreLife-BE/internal/middleware"
	NewError "github.com/AkbarFikri/PreLife-BE/internal/pkg/error"
	"github.com/AkbarFikri/PreLife-BE/internal/pkg/helper"
	"github.com/AkbarFikri/PreLife-BE/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"net/http"
	"time"
)

type ChatbotHandler struct {
	ChatbotService chatbotService.ChatbotService
	Middleware     middleware.Middleware
}

func New(cs chatbotService.ChatbotService, mid middleware.Middleware) *ChatbotHandler {
	return &ChatbotHandler{
		Middleware:     mid,
		ChatbotService: cs,
	}
}

func (h *ChatbotHandler) Endpoints(s *gin.RouterGroup) {
	chat := s.Group("/chatbot", h.Middleware.AuthJWT())
	chat.POST("/chat", h.chat)
	chat.GET("/history", h.history)
}

func (h *ChatbotHandler) chat(ctx *gin.Context) {
	c, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	user := helper.GetUserLoginData(ctx)

	var req dto.ChatRequest
	if err := ctx.ShouldBind(&req); err != nil {
		err = NewError.ErrorBadRequest
		response.New(response.WithError(err)).Send(ctx)
		return
	}

	payload, err := h.ChatbotService.GenerateChatbotResponseFrom(c, user, req.Message)
	if err != nil {
		response.New(response.WithError(err)).Send(ctx)
		return
	}

	response.New(
		response.WithPayload(payload),
		response.WithMessage("successfully get response from gemini"),
		response.WithHttpCode(http.StatusOK),
	).Send(ctx)
	return
}

func (h *ChatbotHandler) history(ctx *gin.Context) {
	c, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	user := helper.GetUserLoginData(ctx)

	payload, err := h.ChatbotService.GetProfileChatbotHistory(c, user)
	if err != nil {
		response.New(response.WithError(err)).Send(ctx)
		return
	}

	response.New(
		response.WithPayload(payload),
		response.WithMessage("successfully get chatbot history"),
		response.WithHttpCode(http.StatusOK),
	).Send(ctx)
}
