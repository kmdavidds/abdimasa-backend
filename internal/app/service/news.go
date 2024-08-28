package service

import (
	"github.com/google/uuid"
	"github.com/kmdavidds/abdimasa-backend/internal/app/repository"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/dto"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/entity"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/validator"
)

type NewsService interface {
	Create(req dto.CreateNewsRequest) error
	GetAll() ([]entity.News, error)
	Update(req dto.UpdateNewsRequest) error
	Delete(req dto.DeleteNewsRequest) error
}

type newsService struct {
	nr repository.NewsRepository
	val validator.Validator
}

func NewNewsService(
	nr repository.NewsRepository,
	val validator.Validator,
) NewsService {
	return &newsService{nr, val}
}

func (ns *newsService) Create(req dto.CreateNewsRequest) error {
	valErr := ns.val.Validate(req)
	if valErr != nil {
		return valErr
	}

	idV7, err := uuid.NewV7()
	if err != nil {
		return err
	}

	news := entity.News{
		ID:          idV7,
		Title:       req.Title,
		Description: req.Description,
		ImageURL:    req.ImageURL,
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

func (ns *newsService) Update(req dto.UpdateNewsRequest) error {
	valErr := ns.val.Validate(req)
	if valErr != nil {
		return valErr
	}

	news := entity.News{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
		ImageURL:    req.ImageURL,
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