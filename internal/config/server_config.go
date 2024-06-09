package config

import (
	userHandler "github.com/AkbarFikri/PreLife-BE/internal/api/user/handler"
	userRepository "github.com/AkbarFikri/PreLife-BE/internal/api/user/repository"
	userService "github.com/AkbarFikri/PreLife-BE/internal/api/user/service"
	"github.com/AkbarFikri/PreLife-BE/internal/middleware"
	"github.com/AkbarFikri/PreLife-BE/internal/pkg/database"
	firebase2 "github.com/AkbarFikri/PreLife-BE/internal/pkg/firebase"
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

	// Firebase utils
	firebaseAuth := firebaseApp.Auth()

	// Middleware
	mid := middleware.New(firebaseAuth, log)

	// Repository init
	usRepository := userRepository.New(db)

	// Service init
	usService := userService.New(usRepository, log)

	// Handler init
	usHandler := userHandler.New(mid, log, usService)

	s.handlers = []Handler{usHandler}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return s
}

func (s *Server) SetupRoute() {
	s.app.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ready!!!"})
	})

	s.app.Use(gin.Logger())

	for _, h := range s.handlers {
		h.Endpoints(s.app.Group("/api/v1"))
	}
}

func (s *Server) Run() {
	s.SetupRoute()
	s.app.Run()
}
