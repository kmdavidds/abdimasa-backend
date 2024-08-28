package service

import (
	"github.com/google/uuid"
	"github.com/kmdavidds/abdimasa-backend/internal/app/repository"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/dto"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/entity"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/errors"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/validator"
)

type ActivityService interface {
	Create(req dto.CreateActivityRequest) error
	GetAll() ([]entity.Activity, error)
	Update(req dto.UpdateActivityRequest) error
	Delete(req dto.DeleteActivityRequest) error
}

type activityService struct {
	ar  repository.ActivityRepository
	val validator.Validator
}

func NewActivityService(
	ar repository.ActivityRepository,
	val validator.Validator,
) ActivityService {
	return &activityService{ar, val}
}

func (as *activityService) Create(req dto.CreateActivityRequest) error {
	valErr := as.val.Validate(req)
	if valErr != nil {
		return valErr
	}

	idV7, err := uuid.NewV7()
	if err != nil {
		return err
	}

	activity := entity.Activity{
		ID:       idV7,
		Title:    req.Title,
		ImageURL: req.ImageURL,
		Date:     req.Date,
		Time:     req.Time,
		Location: req.Location,
	}

	_, err = as.ar.Create(&activity)
	if err != nil {
		return err
	}

	return nil
}

func (as *activityService) GetAll() ([]entity.Activity, error) {
	activities := []entity.Activity{}
	_, err := as.ar.GetAll(&activities)
	if err != nil {
		return nil, err
	}

	return activities, nil
}

func (as *activityService) Update(req dto.UpdateActivityRequest) error {
	valErr := as.val.Validate(req)
	if valErr != nil {
		return valErr
	}

	activity := entity.Activity{
		ID:       req.ID,
		Title:    req.Title,
		ImageURL: req.ImageURL,
		Date:     req.Date,
		Time:     req.Time,
		Location: req.Location,
	}

	count, err := as.ar.Update(&activity)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.ErrorNotFound
	}

	return nil
}

func (as *activityService) Delete(req dto.DeleteActivityRequest) error {
	valErr := as.val.Validate(req)
	if valErr != nil {
		return valErr
	}

	count, err := as.ar.Delete(req.ID)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.ErrorNotFound
	}

	return nil
}
