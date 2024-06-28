package nutritionHandler

import (
	nutritionsService "github.com/AkbarFikri/PreLife-BE/internal/api/nutritions/service"
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

type NutritionHandler struct {
	NutritionService nutritionsService.NutritionService
	Middleware       middleware.Middleware
}

func New(ns nutritionsService.NutritionService, mid middleware.Middleware) *NutritionHandler {
	return &NutritionHandler{
		NutritionService: ns,
		Middleware:       mid,
	}
}

func (h *NutritionHandler) Endpoints(s *gin.RouterGroup) {
	nutritions := s.Group("/nutritions", h.Middleware.AuthJWT())
	nutritions.POST("/generate", h.generateNutritionsPredict)
	nutritions.POST("/", h.storeNutritions)
	nutritions.GET("/", h.currentNutritions)
}

func (h *NutritionHandler) generateNutritionsPredict(ctx *gin.Context) {
	c, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	var req dto.GenerateNutritionRequest
	if err := ctx.ShouldBind(&req); err != nil {
		err = NewError.ErrorBadRequest
		response.New(response.WithError(err)).Send(ctx)
		return
	}

	user := helper.GetUserLoginData(ctx)

	payload, err := h.NutritionService.GeneratePredict(c, user, req)
	if err != nil {
		response.New(response.WithError(err)).SendAbort(ctx)
		return
	}

	response.New(
		response.WithPayload(payload),
		response.WithMessage("successfully generated predict"),
		response.WithHttpCode(http.StatusOK),
	).Send(ctx)
}

func (h *NutritionHandler) storeNutritions(ctx *gin.Context) {
	c, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	var req dto.NutritionsRequest
	if err := ctx.ShouldBind(&req); err != nil {
		err = NewError.ErrorBadRequest
		response.New(response.WithError(err)).Send(ctx)
		return
	}

	user := helper.GetUserLoginData(ctx)

	payload, err := h.NutritionService.StoreNutritions(c, user, req)
	if err != nil {
		response.New(response.WithError(err)).Send(ctx)
		return
	}

	response.New(
		response.WithPayload(payload),
		response.WithMessage("successfully stored profile nutritions"),
		response.WithHttpCode(http.StatusCreated),
	).Send(ctx)
}

func (h *NutritionHandler) currentNutritions(ctx *gin.Context) {
	c, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	user := helper.GetUserLoginData(ctx)

	payload, err := h.NutritionService.CurrentNutritions(c, user)
	if err != nil {
		response.New(response.WithError(err)).Send(ctx)
		return
	}

	response.New(
		response.WithPayload(payload),
		response.WithMessage("successfully calculate current user nutritions"),
		response.WithHttpCode(http.StatusOK),
	).Send(ctx)
}
