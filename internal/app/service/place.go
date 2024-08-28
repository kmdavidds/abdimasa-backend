package service

import (
	"github.com/google/uuid"
	"github.com/kmdavidds/abdimasa-backend/internal/app/repository"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/dto"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/entity"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/errors"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/validator"
)

type PlaceService interface {
	Create(req dto.CreatePlaceRequest) error
	GetAll() ([]entity.Place, error)
	GetByID(req dto.GetPlaceByIDRequest) (entity.Place, error)
	Update(req dto.UpdatePlaceRequest) error
	Delete(req dto.DeletePlaceRequest) error
}

type placeService struct {
	pr  repository.PlaceRepository
	val validator.Validator
}

func NewPlaceService(
	pr repository.PlaceRepository,
	val validator.Validator,
) PlaceService {
	return &placeService{pr, val}
}

func (ps *placeService) Create(req dto.CreatePlaceRequest) error {
	valErr := ps.val.Validate(req)
	if valErr != nil {
		return valErr
	}

	idV7, err := uuid.NewV7()
	if err != nil {
		return err
	}

	place := entity.Place{
		ID: 		idV7,
		Name: 		req.Name,
		Location: 	req.Location,
		Description: req.Description,
		Address: 	req.Address,
		OpeningHours: req.OpeningHours,
		ClosingHours: req.ClosingHours,
		EntryPrice: req.EntryPrice,
		ImageURL1: req.ImageURL1,
		ImageURL2: req.ImageURL2,
		ImageURL3: req.ImageURL3,
		MapURL: req.MapURL,
		Rating: req.Rating,
	}

	_, err = ps.pr.Create(&place)
	if err != nil {
		return err
	}

	return nil
}

func (ps *placeService) GetAll() ([]entity.Place, error) {
	places := []entity.Place{}
	_, err := ps.pr.GetAll(&places)
	if err != nil {
		return nil, err
	}

	return places, nil
}

func (ps *placeService) GetByID(req dto.GetPlaceByIDRequest) (entity.Place, error) {
	valErr := ps.val.Validate(req)
	if valErr != nil {
		return entity.Place{}, valErr
	}

	place := entity.Place{ID: req.ID}

	rowsAffected, err := ps.pr.GetByID(&place)
	if err != nil {
		return entity.Place{}, err
	}

	if rowsAffected == 0 {
		return entity.Place{}, errors.ErrorNotFound
	}

	return place, nil
}

func (ps *placeService) Update(req dto.UpdatePlaceRequest) error {
	valErr := ps.val.Validate(req)
	if valErr != nil {
		return valErr
	}

	place := entity.Place{
		ID: req.ID,
		Name: req.Name,
		Location: req.Location,
		Description: req.Description,
		Address: req.Address,
		OpeningHours: req.OpeningHours,
		ClosingHours: req.ClosingHours,
		EntryPrice: req.EntryPrice,
		ImageURL1: req.ImageURL1,
		ImageURL2: req.ImageURL2,
		ImageURL3: req.ImageURL3,
		MapURL: req.MapURL,
		Rating: req.Rating,
	}

	rowsAffected, err := ps.pr.Update(&place)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.ErrorNotFound
	}

	return nil
}

func (ps *placeService) Delete(req dto.DeletePlaceRequest) error {
	valErr := ps.val.Validate(req)
	if valErr != nil {
		return valErr
	}

	rowsAffected, err := ps.pr.Delete(req.ID)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.ErrorNotFound
	}

	return nil
}