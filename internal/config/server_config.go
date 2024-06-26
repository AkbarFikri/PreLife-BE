package config

import (
	authHandler "github.com/AkbarFikri/PreLife-BE/internal/api/authentication/handler"
	authRepository "github.com/AkbarFikri/PreLife-BE/internal/api/authentication/repository"
	authService "github.com/AkbarFikri/PreLife-BE/internal/api/authentication/service"
	chatbotHandler "github.com/AkbarFikri/PreLife-BE/internal/api/chatbot/handler"
	chatbotRepository "github.com/AkbarFikri/PreLife-BE/internal/api/chatbot/repository"
	chatbotService "github.com/AkbarFikri/PreLife-BE/internal/api/chatbot/service"
	nutritionHandler "github.com/AkbarFikri/PreLife-BE/internal/api/nutritions/handler"
	nutritionRepository "github.com/AkbarFikri/PreLife-BE/internal/api/nutritions/repository"
	nutritionsService "github.com/AkbarFikri/PreLife-BE/internal/api/nutritions/service"
	userHandler "github.com/AkbarFikri/PreLife-BE/internal/api/user/handler"
	userRepository "github.com/AkbarFikri/PreLife-BE/internal/api/user/repository"
	userService "github.com/AkbarFikri/PreLife-BE/internal/api/user/service"
	"github.com/AkbarFikri/PreLife-BE/internal/middleware"
	"github.com/AkbarFikri/PreLife-BE/internal/pkg/database"
	firebase2 "github.com/AkbarFikri/PreLife-BE/internal/pkg/firebase"
	"github.com/AkbarFikri/PreLife-BE/internal/pkg/gemini"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	app      *gin.Engine
	firebase *firebase2.FirebaseClient
	handlers []Handler
	log      *logrus.Logger
}

type Handler interface {
	Endpoints(s *gin.RouterGroup)
}

func NewServer(app *gin.Engine, firebaseApp *firebase2.FirebaseClient, log *logrus.Logger) *Server {
	s := &Server{app: app, log: log}
	db, err := database.NewPostgres()
	if err != nil {
		log.Fatal("Unable connect to database")
	}

	// Third Party
	firebaseAuth := firebaseApp.Auth()
	firebaseStorage := firebaseApp.Storage()
	geminiAi := gemini.New()

	// Middleware
	mid := middleware.New(firebaseAuth, log)

	// Repository init
	usRepository := userRepository.New(db)
	atRepository := authRepository.New(db)
	cbRepository := chatbotRepository.New(db)
	ntRepository := nutritionRepository.New(db)

	// Service init
	usService := userService.New(usRepository, log, firebaseAuth)
	atService := authService.New(atRepository, log, firebaseAuth)
	cbService := chatbotService.New(log, geminiAi, cbRepository)
	ntService := nutritionsService.New(log, ntRepository, geminiAi, firebaseStorage)

	// Handler init
	atHandler := authHandler.New(atService)
	usHandler := userHandler.New(mid, log, usService)
	cbHandler := chatbotHandler.New(cbService, mid)
	ntHandler := nutritionHandler.New(ntService, mid)

	s.handlers = []Handler{usHandler, atHandler, cbHandler, ntHandler}

	return s
}

func (s *Server) SetupRoute() {
	s.app.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ready!!!"})
	})

	s.app.Use(gin.Logger())
	s.app.Use(middleware.CORSMiddleware())

	for _, h := range s.handlers {
		h.Endpoints(s.app.Group("/api/v1"))
	}
}

func (s *Server) Run() {
	s.SetupRoute()
	s.app.Run()
}
