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

type NewsService interface {
	Create(req dto.CreateNewsRequest) error
	GetAll() ([]entity.News, error)
	GetByID(req dto.GetNewsByIDRequest) (entity.News, error)
	Update(req dto.UpdateNewsRequest) error
	Delete(req dto.DeleteNewsRequest) error
}

type newsService struct {
	nr       repository.NewsRepository
	val      validator.Validator
	supabase supabase.Supabase
}

func NewNewsService(
	nr repository.NewsRepository,
	val validator.Validator,
	supabase supabase.Supabase,
) NewsService {
	return &newsService{nr, val, supabase}
}

func (ns *newsService) Create(req dto.CreateNewsRequest) error {
	valErr := ns.val.Validate(req)
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

		imageURL, err = ns.supabase.Upload(req.Image1)
		if err != nil {
			return err
		}
	}	

	idV7, err := uuid.NewV7()
	if err != nil {
		return err
	}

	news := entity.News{
		ID:          idV7,
		Title:       req.Title,
		Description: req.Description,
		ImageURL:    imageURL,
	}

	_, err = ns.nr.Create(&news)
	if err != nil {
		return err
	}

	return nil
}

func (ns *newsService) GetAll() ([]entity.News, error) {
	news := []entity.News{}
	_, err := ns.nr.GetAll(&news)
	if err != nil {
		return nil, err
	}

	return news, nil
}

func (ns *newsService) GetByID(req dto.GetNewsByIDRequest) (entity.News, error) {
	valErr := ns.val.Validate(req)
	if valErr != nil {
		return entity.News{}, valErr
	}

	news := entity.News{ID: req.ID}

	rowsAffected, err := ns.nr.GetByID(&news)
	if err != nil {
		return entity.News{}, err
	}

	if rowsAffected == 0 {
		return entity.News{}, errors.ErrorNotFound
	}

	return news, nil
}

func (ns *newsService) Update(req dto.UpdateNewsRequest) error {
	valErr := ns.val.Validate(req)
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

		news := entity.News{ID: req.ID}
		rowsAffected, err := ns.nr.GetByID(&news)
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return errors.ErrorNotFound
		}

		ns.supabase.Delete(news.ImageURL)

		imageURL, err = ns.supabase.Upload(req.Image1)
		if err != nil {
			return err
		}
	}	

	news := entity.News{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
		ImageURL:    imageURL,
	}

	_, err := ns.nr.Update(&news)
	if err != nil {
		return err
	}

	return nil
}

func (ns *newsService) Delete(req dto.DeleteNewsRequest) error {
	_, err := ns.nr.Delete(req.ID)
	if err != nil {
		return err
	}

	return nil
}
