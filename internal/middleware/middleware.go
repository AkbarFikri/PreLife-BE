package middleware

import (
	"firebase.google.com/go/v4/auth"
	"github.com/sirupsen/logrus"
)

type Middleware struct {
	authClient *auth.Client
	log        *logrus.Logger
}

func New(client *auth.Client, log *logrus.Logger) Middleware {
	return Middleware{
		authClient: client,
		log:        log,
	}
}
