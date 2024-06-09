package userService

import (
	userRepository "github.com/AkbarFikri/PreLife-BE/internal/api/user/repository"
	"github.com/AkbarFikri/PreLife-BE/internal/domain"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type UserService interface {
	SaveUserIfNotExists(ctx context.Context, user domain.User) (domain.User, error)
}

type userService struct {
	UserRepository userRepository.UserRepository
	log            *logrus.Logger
}

func New(userRepository userRepository.UserRepository, log *logrus.Logger) UserService {
	return &userService{
		UserRepository: userRepository,
		log:            log,
	}
}

func (s userService) SaveUserIfNotExists(ctx context.Context, user domain.User) (domain.User, error) {
	count, err := s.UserRepository.CountEmail(ctx, user.Email)
	if err != nil {
		s.log.Errorf("UserRepository.CountEmail err: %v", err)
		return domain.User{}, err
	}

	if count > 0 {
		user, err := s.UserRepository.FindById(ctx, user.ID)
		if err != nil {
			s.log.Errorf("unable to find user by id: %v", err)
			return user, err
		}
		return user, nil
	}

	if err := s.UserRepository.Save(ctx, user); err != nil {
		s.log.Errorf("error when save user: %v", err)
		return domain.User{}, err
	}

	return user, nil
}
