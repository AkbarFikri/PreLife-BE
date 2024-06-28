package nutritionsService

import (
	"database/sql"
	"errors"
	"fmt"
	nutritionRepository "github.com/AkbarFikri/PreLife-BE/internal/api/nutritions/repository"
	"github.com/AkbarFikri/PreLife-BE/internal/domain"
	"github.com/AkbarFikri/PreLife-BE/internal/dto"
	"github.com/AkbarFikri/PreLife-BE/internal/pkg/firebase"
	"github.com/AkbarFikri/PreLife-BE/internal/pkg/gemini"
	"github.com/AkbarFikri/PreLife-BE/internal/pkg/helper"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"io"
	"time"
)

type NutritionService interface {
	GeneratePredict(context.Context, dto.UserTokenData, dto.GenerateNutritionRequest) (dto.GenerateNutritionsResponse, error)
	StoreNutritions(context.Context, dto.UserTokenData, dto.NutritionsRequest) (dto.NutritionsResponse, error)
	CurrentNutritions(ctx context.Context, data dto.UserTokenData) (dto.CurrentNutritionsResponse, error)
}

type nutritionService struct {
	log                 *logrus.Logger
	NutritionRepository nutritionRepository.NutritionRepository
	gemini              *gemini.Gemini
	firebaseStorage     firebase.FirebaseStorage
}

func New(log *logrus.Logger, nr nutritionRepository.NutritionRepository, ai *gemini.Gemini, fs firebase.FirebaseStorage) NutritionService {
	return &nutritionService{
		NutritionRepository: nr,
		gemini:              ai,
		log:                 log,
		firebaseStorage:     fs,
	}
}

func (s nutritionService) GeneratePredict(c context.Context, user dto.UserTokenData, request dto.GenerateNutritionRequest) (dto.GenerateNutritionsResponse, error) {
	pict, err := request.Picture.Open()
	if err != nil {
		s.log.Errorf("error when opening picture from user : %v", err)
		return dto.GenerateNutritionsResponse{}, err
	}
	defer pict.Close()

	fileBytes, err := io.ReadAll(pict)
	if err != nil {
		s.log.Errorf("error when reading picture from user : %v", err)
		return dto.GenerateNutritionsResponse{}, err
	}

	predict, err := s.gemini.PredictFoodNutritions(c, fileBytes)
	if err != nil {
		s.log.Errorf("error when gemini.PredictFoodNutritions %v", err)
		return dto.GenerateNutritionsResponse{}, err
	}

	if predict.FoodName == "notfood" {
		return dto.GenerateNutritionsResponse{}, ErrorInvalidPictureProvided
	}

	return predict, nil
}

func (s nutritionService) StoreNutritions(c context.Context, user dto.UserTokenData, request dto.NutritionsRequest) (dto.NutritionsResponse, error) {
	pict, err := request.Picture.Open()
	if err != nil {
		s.log.Errorf("error when opening picture from user : %v", err)
		return dto.NutritionsResponse{}, err
	}
	defer pict.Close()

	fileBytes, err := io.ReadAll(pict)
	if err != nil {
		s.log.Errorf("error when reading picture from user : %v", err)
		return dto.NutritionsResponse{}, err
	}

	generateFileName := fmt.Sprintf("nutritions-%s-%d", user.ProfileId, time.Now().UnixNano())

	link, err := s.firebaseStorage.UploadFile(fileBytes, generateFileName)
	if err != nil {
		s.log.Errorf("error when uploading file : %v", err)
		return dto.NutritionsResponse{}, err
	}

	nutritions := domain.Nutrition{
		ID:            fmt.Sprintf("nutri-%s", helper.GenerateUID(16)),
		UserProfileId: user.ProfileId,
		FoodName:      request.FoodName,
		Calories:      request.Calories,
		Carbohydrate:  request.Carbohidrate,
		Protein:       request.Protein,
		PictureLink:   link,
		CreatedAt:     time.Now(),
	}

	if err := s.NutritionRepository.Save(c, nutritions); err != nil {
		s.log.Errorf("error when saving nutritions : %v", err)
		return dto.NutritionsResponse{}, err
	}

	return dto.NutritionsResponse{
		ID:          nutritions.ID,
		PictureLink: nutritions.PictureLink,
		FoodName:    nutritions.FoodName,
	}, nil
}

func (s nutritionService) CurrentNutritions(ctx context.Context, request dto.UserTokenData) (dto.CurrentNutritionsResponse, error) {
	nutritionsHistory, err := s.NutritionRepository.GetCurrentDateHistory(ctx, request.ProfileId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
		} else {
			s.log.Errorf("error when getting current nutritions history : %v", err)
			return dto.CurrentNutritionsResponse{}, err
		}
	}

	var res dto.CurrentNutritionsResponse

	if len(nutritionsHistory) == 0 {
		res.CurrentCarbohydrates = 0
		res.CurrentCalories = 0
		res.CurrentProtein = 0
	} else {
		for _, nutrition := range nutritionsHistory {
			res.CurrentCarbohydrates += nutrition.Carbohydrate
			res.CurrentCalories += nutrition.Calories
			res.CurrentProtein += nutrition.Protein
		}
	}

	switch request.ProfileType {
	case PregnantProfile:
		res.MaxCalories = 2550
		res.MaxCarbohydrates = 350
		res.MaxProtein = 100
	case NonPregnantProfile:
		res.MaxCalories = 2550
		res.MaxCarbohydrates = 350
		res.MaxProtein = 100
	}

	return res, nil
}
