package authHandler

import (
	authService "github.com/AkbarFikri/PreLife-BE/internal/api/authentication/service"
	"github.com/AkbarFikri/PreLife-BE/internal/dto"
	NewError "github.com/AkbarFikri/PreLife-BE/internal/pkg/error"
	"github.com/AkbarFikri/PreLife-BE/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"net/http"
	"time"
)

type AuthHandler struct {
	AuthService authService.AuthService
}

func New(as authService.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: as,
	}
}

func (h *AuthHandler) Endpoints(s *gin.RouterGroup) {
	auth := s.Group("/auth")
	auth.POST("/register", h.Register)
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	var req dto.AuthRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		err = NewError.ErrorBadRequest
		response.New(response.WithError(err)).Send(ctx)
		return
	}

	payload, err := h.AuthService.Register(c, req)
	if err != nil {
		response.New(response.WithError(err)).Send(ctx)
		return
	}

	response.New(
		response.WithPayload(payload),
		response.WithMessage("successfully registered new user"),
		response.WithHttpCode(http.StatusCreated),
	).Send(ctx)
	return
}
