package authService

import (
	"firebase.google.com/go/v4/auth"
	"github.com/AkbarFikri/PreLife-BE/internal/api/authentication/repository"
	"github.com/AkbarFikri/PreLife-BE/internal/domain"
	"github.com/AkbarFikri/PreLife-BE/internal/dto"
	NewError "github.com/AkbarFikri/PreLife-BE/internal/pkg/error"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"time"
)

type AuthService interface {
	Register(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error)
}

type authService struct {
	log            *logrus.Logger
	AuthRepository authRepository.AuthRepository
	authClient     *auth.Client
}

func New(ar authRepository.AuthRepository, log *logrus.Logger, client *auth.Client) AuthService {
	return &authService{
		AuthRepository: ar,
		authClient:     client,
		log:            log,
	}
}

func (s authService) Register(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error) {
	count, err := s.AuthRepository.CountEmail(ctx, req.Email)
	if err != nil {
		s.log.Errorf("error repository not recognized %v", err)
		return dto.AuthResponse{}, NewError.ErrorGeneral
	}

	if count > 0 {
		s.log.Printf("error email already exists")
		return dto.AuthResponse{}, ErrorEmailAlreadyUsed
	}

	dateBirth, err := time.Parse("02-01-2006", req.DateOfBirth)
	if err != nil {
		s.log.Errorf("error date format %v", err)
		return dto.AuthResponse{}, ErrorInvalidDateFormat
	}

	user := domain.User{
		Email:       req.Email,
		FullName:    req.FullName,
		DateOfBirth: dateBirth,
	}

	resClient, err := s.authClient.CreateUser(ctx, (&auth.UserToCreate{}).DisplayName(user.FullName).Email(req.Email).Password(req.Password))
	if err != nil {
		s.log.Errorf("error create user on firebase : %v", err)
		if auth.IsEmailAlreadyExists(err) {
			return dto.AuthResponse{}, ErrorEmailAlreadyUsed
		}
		return dto.AuthResponse{}, err
	}

	user.ID = resClient.UID

	if err := s.AuthRepository.Save(ctx, user); err != nil {
		s.log.Errorf("error repository when create user : %v", err)
		return dto.AuthResponse{}, NewError.ErrorGeneral
	}

	return dto.AuthResponse{
		FullName:    user.FullName,
		Email:       user.Email,
		ID:          user.ID,
		DateOfBirth: user.DateOfBirth.Format("02-01-2006"),
	}, nil
}
