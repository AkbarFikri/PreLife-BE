package userHandler

import (
	"fmt"
	userService "github.com/AkbarFikri/PreLife-BE/internal/api/user/service"
	"github.com/AkbarFikri/PreLife-BE/internal/dto"
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
	middleware  middleware.Middleware
	log         *logrus.Logger
	userService userService.UserService
}

func New(middleware middleware.Middleware,
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
	user.GET("/current", h.middleware.AuthJWT(), h.current)
	user.GET("/current/profile", h.middleware.AuthJWT(), h.currentProfile)
	user.GET("/profiles", h.middleware.AuthJWT(), h.profiles)
	user.POST("/profile/new", h.middleware.ApiKey(), h.createFreshProfile)
}

func (h UserHandler) current(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	user := helper.GetUserLoginData(ctx)

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

func (h UserHandler) currentProfile(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	user := helper.GetUserLoginData(ctx)

	fmt.Println(user.ProfileType)
	switch user.ProfileType {
	case PregnantProfile:
		payload, err := h.userService.GetPregnantProfileDetails(c, user)
		if err != nil {
			response.New(response.WithError(err)).Send(ctx)
			return
		}

		fmt.Println("masuk")
		response.New(
			response.WithPayload(payload),
			response.WithMessage("successfully find user profile"),
			response.WithHttpCode(http.StatusOK),
		).Send(ctx)
	case NonPregnantProfile:
		payload, err := h.userService.GetNonPregnantProfileDetails(c, user)
		if err != nil {
			response.New(response.WithError(err)).Send(ctx)
			return
		}

		response.New(
			response.WithPayload(payload),
			response.WithMessage("successfully find user profile"),
			response.WithHttpCode(http.StatusOK),
		).Send(ctx)
	}
}

func (h UserHandler) profiles(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	user := helper.GetUserLoginData(ctx)

	payload, err := h.userService.GetUserProfiles(c, user)
	if err != nil {
		response.New(response.WithError(err)).Send(ctx)
		return
	}

	response.New(
		response.WithPayload(payload),
		response.WithMessage("successfully find user profiles"),
		response.WithHttpCode(http.StatusOK),
	).Send(ctx)
}

func (h UserHandler) createFreshProfile(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	var check dto.CheckIsPregnant
	if err := ctx.ShouldBindQuery(&check); err != nil {
		h.log.Errorf("bad request: %v", err)
		err = NewError.ErrorBadRequest
		response.New(response.WithError(err)).Send(ctx)
		return
	}

	switch check.ProfileType {
	case PregnantProfile:
		var req dto.RegisterProfilePregnantRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			h.log.Errorf("bad request: %v", req)
			err = NewError.ErrorBadRequest
			response.New(response.WithError(err), response.WithPayload(req)).Send(ctx)
			return
		}

		req.IsPregnant = true
		payload, err := h.userService.RegisterPregnantProfile(c, req)
		if err != nil {
			response.New(response.WithError(err)).SendAbort(ctx)
			return
		}

		response.New(
			response.WithPayload(payload),
			response.WithMessage("succesfully create fresh profile pregnant for user"),
			response.WithHttpCode(http.StatusCreated),
		).Send(ctx)
	case NonPregnantProfile:
		var req dto.RegisterProfileNotPregnantRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			h.log.Errorf("bad request: %v", req)
			err = NewError.ErrorBadRequest
			response.New(response.WithError(err), response.WithPayload(req)).Send(ctx)
			return
		}

		req.IsPregnant = false
		payload, err := h.userService.RegisterNotPregnantProfile(c, req)
		if err != nil {
			response.New(response.WithError(err)).Send(ctx)
			return
		}

		response.New(
			response.WithPayload(payload),
			response.WithMessage("succesfully create fresh profile non pregnant for user"),
			response.WithHttpCode(http.StatusCreated),
		).Send(ctx)
	}

}
