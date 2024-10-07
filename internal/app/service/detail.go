package service

import (
	"github.com/kmdavidds/abdimasa-backend/internal/app/repository"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/dto"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/entity"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/errors"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/validator"
)

type DetailService interface {
	Create(req dto.CreateDetailRequest) error
	GetAll() ([]entity.Detail, error)
	GetByID(req dto.GetDetailByIDRequest) (entity.Detail, error)
	GetBySlug(req dto.GetDetailBySlugRequest) (entity.Detail, error)
	Update(req dto.UpdateDetailRequest) error
	Delete(req dto.DeleteDetailRequest) error
}

type detailService struct {
	dr  repository.DetailRepository
	val validator.Validator
}

func NewDetailService(
	dr repository.DetailRepository,
	val validator.Validator,
) DetailService {
	return &detailService{dr, val}
}

func (ds *detailService) Create(req dto.CreateDetailRequest) error {
	valErr := ds.val.Validate(req)
	if valErr != nil {
		return valErr
	}

	detail := entity.Detail{
		Slug: req.Slug,
		Value: req.Value,
	}

	_, err := ds.dr.Create(&detail)
	if err != nil {
		return err
	}

	return nil
}

func (ds *detailService) GetAll() ([]entity.Detail, error) {
	detail := []entity.Detail{}
	_, err := ds.dr.GetAll(&detail)
	if err != nil {
		return nil, err
	}

	return detail, nil
}

func (ds *detailService) GetByID(req dto.GetDetailByIDRequest) (entity.Detail, error) {
	valErr := ds.val.Validate(req)
	if valErr != nil {
		return entity.Detail{}, valErr
	}

	detail := entity.Detail{ID: req.ID}

	rowsAffected, err := ds.dr.GetByID(&detail)
	if err != nil {
		return entity.Detail{}, err
	}

	if rowsAffected == 0 {
		return entity.Detail{}, errors.ErrorNotFound
	}

	return detail, nil
}

func (ds *detailService) GetBySlug(req dto.GetDetailBySlugRequest) (entity.Detail, error) {
	valErr := ds.val.Validate(req)
	if valErr != nil {
		return entity.Detail{}, valErr
	}

	detail := entity.Detail{Slug: req.Slug}

	rowsAffected, err := ds.dr.GetBySlug(&detail)
	if err != nil {
		return entity.Detail{}, err
	}

	if rowsAffected == 0 {
		return entity.Detail{}, errors.ErrorNotFound
	}

	return detail, nil
}

func (ds *detailService) Update(req dto.UpdateDetailRequest) error {
	valErr := ds.val.Validate(req)
	if valErr != nil {
		return valErr
	}

	detail := entity.Detail{
		ID: req.ID,
		Slug: req.Slug,
		Value: req.Value,
	}

	_, err := ds.dr.Update(&detail)
	if err != nil {
		return err
	}

	return nil
}

func (ds *detailService) Delete(req dto.DeleteDetailRequest) error {
	_, err := ds.dr.Delete(req.ID)
	if err != nil {
		return err
	}

	return nil
}
