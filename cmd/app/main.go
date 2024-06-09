package main

import (
	"github.com/AkbarFikri/PreLife-BE/internal/config"
	firebase2 "github.com/AkbarFikri/PreLife-BE/internal/pkg/firebase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load()
	env := os.Getenv("ENV")
	if err != nil && env == "" {
		panic("Error loading .env file")
	}

	logger := config.NewLogger()
	firebaseClient, err := firebase2.Init(logger)
	if err != nil {
		panic(err)
	}
	gin := gin.Default()
	app := config.NewServer(gin, firebaseClient, logger)

	app.Run()
}
