package service

import (
	"slices"

	"github.com/google/uuid"
	"github.com/kmdavidds/abdimasa-backend/internal/app/repository"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/dto"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/entity"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/errors"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/fileops"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/supabase"
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
	supabase supabase.Supabase
}

func NewActivityService(
	ar repository.ActivityRepository,
	val validator.Validator,
	supabase supabase.Supabase,
) ActivityService {
	return &activityService{ar, val, supabase}
}

func (as *activityService) Create(req dto.CreateActivityRequest) error {
	valErr := as.val.Validate(req)
	if valErr != nil {
		return valErr
	}

	var imageURL = ""

	if req.Image1 != nil {
		if req.Image1.Size > 1*fileops.MegaByte {
			return errors.ErrorFileTooLarge
		}

		fileType, err := fileops.DetectMultipartFileType(req.Image1)

		if err != nil {
			return errors.ErrorInvalidFileType
		}

		allowedTypes := fileops.ImageContentTypes
		if !slices.Contains(allowedTypes, fileType) {
			return errors.ErrorInvalidFileType
		}

		imageURL, err = as.supabase.Upload(req.Image1)
		if err != nil {
			return err
		}
	}	

	idV7, err := uuid.NewV7()
	if err != nil {
		return err
	}

	activity := entity.Activity{
		ID:       idV7,
		Title:    req.Title,
		ImageURL: imageURL,
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

	var imageURL = ""

	if req.Image1 != nil {
		if req.Image1.Size > 1*fileops.MegaByte {
			return errors.ErrorFileTooLarge
		}

		fileType, err := fileops.DetectMultipartFileType(req.Image1)

		if err != nil {
			return errors.ErrorInvalidFileType
		}

		allowedTypes := fileops.ImageContentTypes
		if !slices.Contains(allowedTypes, fileType) {
			return errors.ErrorInvalidFileType
		}

		imageURL, err = as.supabase.Upload(req.Image1)
		if err != nil {
			return err
		}
	}	

	activity := entity.Activity{
		ID:       req.ID,
		Title:    req.Title,
		ImageURL: imageURL,
		Date:     req.Date,
		Time:     req.Time,
		Location: req.Location,
	}

	rowsAffected, err := as.ar.Update(&activity)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.ErrorNotFound
	}

	return nil
}

func (as *activityService) Delete(req dto.DeleteActivityRequest) error {
	valErr := as.val.Validate(req)
	if valErr != nil {
		return valErr
	}

	rowsAffected, err := as.ar.Delete(req.ID)
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		return errors.ErrorNotFound
	}

	return nil
}
