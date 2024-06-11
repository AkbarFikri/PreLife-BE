package userService

import (
	"firebase.google.com/go/v4/auth"
	userRepository "github.com/AkbarFikri/PreLife-BE/internal/api/user/repository"
	"github.com/AkbarFikri/PreLife-BE/internal/domain"
	"github.com/AkbarFikri/PreLife-BE/internal/dto"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"time"
)

type UserService interface {
	SaveUserIfNotExists(ctx context.Context, user domain.User) (dto.AuthResponse, error)
}

type userService struct {
	UserRepository userRepository.UserRepository
	authClient     *auth.Client
	log            *logrus.Logger
}

func New(userRepository userRepository.UserRepository, log *logrus.Logger, client *auth.Client) UserService {
	return &userService{
		UserRepository: userRepository,
		authClient:     client,
		log:            log,
	}
}

func (s userService) SaveUserIfNotExists(ctx context.Context, user domain.User) (dto.AuthResponse, error) {
	count, err := s.UserRepository.CountEmail(ctx, user.Email)
	if err != nil {
		s.log.Errorf("UserRepository.CountEmail err: %v", err)
		return dto.AuthResponse{}, err
	}

	if count > 0 {
		user, err := s.UserRepository.FindUserByEmail(ctx, user.Email)
		if err != nil {
			s.log.Errorf("unable to find user by id: %v", err)
			return dto.AuthResponse{}, err
		}

		res := dto.AuthResponse{
			ID:          user.ID,
			Email:       user.Email,
			FullName:    user.FullName,
			DateOfBirth: user.DateOfBirth.Format("02-01-2006"),
		}

		return res, nil
	}

	firebaseUser, err := s.authClient.GetUser(ctx, user.ID)
	if err != nil {
		s.log.Errorf("unable to find user by id in firebase: %v", err)
		return dto.AuthResponse{}, err
	}

	user.FullName = firebaseUser.DisplayName
	if user.FullName == "" {
		user.FullName = "default-name"
	}
	user.DateOfBirth = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)

	if err := s.UserRepository.Save(ctx, user); err != nil {
		s.log.Errorf("error when save user: %v", err)
		return dto.AuthResponse{}, err
	}

	return dto.AuthResponse{
		ID:          user.ID,
		Email:       user.Email,
		FullName:    user.FullName,
		DateOfBirth: user.DateOfBirth.Format("02-01-2006"),
	}, nil
}
