package userHandler

import (
	userService "github.com/AkbarFikri/PreLife-BE/internal/api/user/service"
	"github.com/AkbarFikri/PreLife-BE/internal/middleware"
	NewError "github.com/AkbarFikri/PreLife-BE/internal/pkg/error"
	"github.com/AkbarFikri/PreLife-BE/internal/pkg/helper"
	"github.com/AkbarFikri/PreLife-BE/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"net/http"
	"time"
)

type UserHandler struct {
	middleware  *middleware.Middleware
	log         *logrus.Logger
	userService userService.UserService
}

func New(middleware *middleware.Middleware,
	log *logrus.Logger,
	us userService.UserService) UserHandler {
	return UserHandler{
		middleware:  middleware,
		log:         log,
		userService: us,
	}
}

func (h UserHandler) Endpoints(s *gin.RouterGroup) {
	user := s.Group("/user")
	user.Use(h.middleware.AuthJWT())
	user.GET("/current", h.current)
}

func (h UserHandler) current(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	user, err := helper.UserDataFromToken(ctx)
	if err != nil {
		err = NewError.ErrorBadRequest
		response.New(response.WithError(err)).Send(ctx)
	}

	payload, err := h.userService.SaveUserIfNotExists(c, user)
	if err != nil {
		response.New(response.WithError(err)).SendAbort(ctx)
		return
	}

	response.New(
		response.WithPayload(payload),
		response.WithMessage("succesfully find/saved user"),
		response.WithHttpCode(http.StatusCreated),
	).Send(ctx)
}
