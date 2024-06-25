package userService

import (
	"database/sql"
	"errors"
	"firebase.google.com/go/v4/auth"
	"fmt"
	userRepository "github.com/AkbarFikri/PreLife-BE/internal/api/user/repository"
	"github.com/AkbarFikri/PreLife-BE/internal/domain"
	"github.com/AkbarFikri/PreLife-BE/internal/dto"
	NewError "github.com/AkbarFikri/PreLife-BE/internal/pkg/error"
	"github.com/AkbarFikri/PreLife-BE/internal/pkg/helper"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"time"
)

type UserService interface {
	SaveUserIfNotExists(ctx context.Context, user dto.UserTokenData) (dto.AuthResponse, error)
	RegisterPregnantProfile(ctx context.Context, req dto.RegisterProfilePregnantRequest) (dto.RegisterProfileResponse, error)
	RegisterNotPregnantProfile(ctx context.Context, req dto.RegisterProfileNotPregnantRequest) (dto.RegisterProfileResponse, error)
	GetPregnantProfileDetails(ctx context.Context, req dto.UserTokenData) (dto.PregnantProfileResponse, error)
	GetNonPregnantProfileDetails(ctx context.Context, req dto.UserTokenData) (dto.NonPregnantProfileResponse, error)
	GetUserProfiles(ctx context.Context, user dto.UserTokenData) ([]dto.ProfileListResponse, error)
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

func (s userService) SaveUserIfNotExists(ctx context.Context, user dto.UserTokenData) (dto.AuthResponse, error) {
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

	var newUser domain.User

	newUser.FullName = firebaseUser.DisplayName
	if newUser.FullName == "" {
		newUser.FullName = "default-name"
	}
	newUser.DateOfBirth = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	newUser.ID = firebaseUser.UID
	newUser.Email = firebaseUser.Email

	if err := s.UserRepository.Save(ctx, newUser); err != nil {
		s.log.Errorf("error when save user: %v", err)
		return dto.AuthResponse{}, err
	}

	return dto.AuthResponse{
		ID:          newUser.ID,
		Email:       newUser.Email,
		FullName:    newUser.FullName,
		DateOfBirth: newUser.DateOfBirth.Format("02-01-2006"),
	}, nil
}

func (s userService) RegisterPregnantProfile(ctx context.Context, req dto.RegisterProfilePregnantRequest) (dto.RegisterProfileResponse, error) {
	user, err := s.UserRepository.FindUserById(ctx, req.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.log.Warnf("Something try to create a new user profile with profile id : %v", req.UserID)
			return dto.RegisterProfileResponse{}, ErrorInvalidUserId
		} else {
			s.log.Errorf("unable to find user by id: %v", err)
			return dto.RegisterProfileResponse{}, NewError.ErrorGeneral
		}
	}

	pregnantDateParse, err := time.Parse("02-01-2006", req.PregnantDate)
	if err != nil {
		s.log.Errorf("unable to parse pregnant date: %v", err)
		return dto.RegisterProfileResponse{}, ErrorInvalidTimeFormat
	}

	profile := domain.PregnantProfile{
		ID:           fmt.Sprintf("pf-p-%s", helper.GenerateUID(20)),
		UserId:       user.ID,
		ProfileName:  user.FullName,
		IsPregnant:   req.IsPregnant,
		PregnantDate: pregnantDateParse,
	}

	claims := map[string]interface{}{
		"profile_id":   profile.ID,
		"id":           user.ID,
		"role_id":      user.RoleId,
		"email":        user.Email,
		"profile_type": 1,
	}

	if err := s.authClient.SetCustomUserClaims(ctx, user.ID, claims); err != nil {
		s.log.Errorf("error when set custom user claims: %v", err)
		return dto.RegisterProfileResponse{}, err
	}

	if err := s.UserRepository.SaveUserPregnantProfile(ctx, profile); err != nil {
		s.log.Errorf("error when save user profile: %v", err)
		return dto.RegisterProfileResponse{}, err
	}

	return dto.RegisterProfileResponse{
		ID:         profile.ID,
		IsPregnant: profile.IsPregnant,
		UserID:     profile.UserId,
	}, nil
}

func (s userService) RegisterNotPregnantProfile(ctx context.Context, req dto.RegisterProfileNotPregnantRequest) (dto.RegisterProfileResponse, error) {
	user, err := s.UserRepository.FindUserById(ctx, req.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.log.Warnf("Something try to create a new user profile with profile id : %v", req.UserID)
			return dto.RegisterProfileResponse{}, ErrorInvalidUserId
		} else {
			s.log.Errorf("unable to find user by id: %v", err)
			return dto.RegisterProfileResponse{}, NewError.ErrorGeneral
		}
	}

	birthDateParse, err := time.Parse("02-01-2006", req.BirthDate)
	if err != nil {
		s.log.Errorf("unable to parse pregnant date: %v", err)
		return dto.RegisterProfileResponse{}, ErrorInvalidTimeFormat
	}

	profile := domain.NotPregnantProfile{
		ID:          fmt.Sprintf("pf-np-%s", helper.GenerateUID(20)),
		UserId:      user.ID,
		ProfileName: user.FullName,
		IsPregnant:  req.IsPregnant,
		BirthDate:   birthDateParse,
		Height:      req.Height,
		Weight:      req.Weight,
	}

	claims := map[string]interface{}{
		"profile_id":   profile.ID,
		"id":           user.ID,
		"role_id":      user.RoleId,
		"email":        user.Email,
		"profile_type": 2,
	}

	if err := s.authClient.SetCustomUserClaims(ctx, user.ID, claims); err != nil {
		s.log.Errorf("error when set custom user claims: %v", err)
		return dto.RegisterProfileResponse{}, err
	}

	if err := s.UserRepository.SaveUserNotPregnantProfile(ctx, profile); err != nil {
		s.log.Errorf("error when save user profile: %v", err)
		return dto.RegisterProfileResponse{}, err
	}

	return dto.RegisterProfileResponse{
		ID:         profile.ID,
		IsPregnant: profile.IsPregnant,
		UserID:     profile.UserId,
	}, nil
}

func (s userService) GetPregnantProfileDetails(ctx context.Context, req dto.UserTokenData) (dto.PregnantProfileResponse, error) {
	profile, err := s.UserRepository.FindPregnantProfileById(ctx, req.ProfileId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return dto.PregnantProfileResponse{}, ErrorInvalidProfileId
		} else {
			s.log.Errorf("unable to find user profile by id: %v", err)
			return dto.PregnantProfileResponse{}, NewError.ErrorGeneral
		}
	}

	now := time.Now()
	timeDiff := now.Sub(profile.PregnantDate)

	pregnantAgeInDay := int(timeDiff.Hours() / 24)
	pregnantAgeInWeek := pregnantAgeInDay / 7
	pregnantAgeInMonth := pregnantAgeInDay / 30

	return dto.PregnantProfileResponse{
		UserID:             profile.UserId,
		PregnantDate:       profile.PregnantDate,
		IsPregnant:         profile.IsPregnant,
		PregnantAgeInDay:   pregnantAgeInDay,
		PregnantAgeInWeek:  pregnantAgeInWeek,
		PregnantAgeInMonth: pregnantAgeInMonth,
		Id:                 profile.ID,
		ProfileName:        profile.ProfileName,
	}, nil
}

func (s userService) GetNonPregnantProfileDetails(ctx context.Context, req dto.UserTokenData) (dto.NonPregnantProfileResponse, error) {
	profile, err := s.UserRepository.FindNonPregnantProfileById(ctx, req.ProfileId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return dto.NonPregnantProfileResponse{}, ErrorInvalidProfileId
		} else {
			s.log.Errorf("unable to find user profile by id: %v", err)
			return dto.NonPregnantProfileResponse{}, NewError.ErrorGeneral
		}
	}

	now := time.Now()
	timeDiff := now.Sub(profile.BirthDate)

	childAgeInDay := int(timeDiff.Hours() / 24)
	childAgeInWeek := childAgeInDay / 7
	childAgeInMonth := childAgeInDay / 30
	childAgeInYear := childAgeInDay / 365

	return dto.NonPregnantProfileResponse{
		UserID:          profile.UserId,
		BirthDate:       profile.BirthDate,
		IsPregnant:      profile.IsPregnant,
		ChildAgeInDay:   childAgeInDay,
		ChildAgeInWeek:  childAgeInWeek,
		ChildAgeInMonth: childAgeInMonth,
		ChildAgeInYear:  childAgeInYear,
		Id:              profile.ID,
		ProfileName:     profile.ProfileName,
		Weight:          profile.Weight,
		Height:          profile.Height,
	}, nil
}

func (s userService) GetUserProfiles(ctx context.Context, user dto.UserTokenData) ([]dto.ProfileListResponse, error) {
	profiles, err := s.UserRepository.FindAllUserProfiles(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	var res []dto.ProfileListResponse

	for _, profile := range profiles {
		response := dto.ProfileListResponse{
			ID:          profile.ID,
			UserID:      profile.UserID,
			IsPregnant:  profile.IsPregnant,
			ProfileName: profile.ProfileName,
		}

		res = append(res, response)
	}

	return res, nil
}
