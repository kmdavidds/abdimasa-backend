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
	supabase supabase.Supabase
}

func NewPlaceService(
	pr repository.PlaceRepository,
	val validator.Validator,
	supabase supabase.Supabase,
) PlaceService {
	return &placeService{pr, val, supabase}
}

func (ps *placeService) Create(req dto.CreatePlaceRequest) error {
	valErr := ps.val.Validate(req)
	if valErr != nil {
		return valErr
	}

	var imageURL1 = ""
	var imageURL2 = ""
	var imageURL3 = ""

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

		imageURL1, err = ps.supabase.Upload(req.Image1)
		if err != nil {
			return err
		}
	}
	if req.Image2 != nil {
		if req.Image2.Size > 1*fileops.MegaByte {
			return errors.ErrorFileTooLarge
		}

		fileType, err := fileops.DetectMultipartFileType(req.Image2)

		if err != nil {
			return errors.ErrorInvalidFileType
		}

		allowedTypes := fileops.ImageContentTypes
		if !slices.Contains(allowedTypes, fileType) {
			return errors.ErrorInvalidFileType
		}

		imageURL2, err = ps.supabase.Upload(req.Image2)
		if err != nil {
			return err
		}
	}
	if req.Image3 != nil {
		if req.Image3.Size > 1*fileops.MegaByte {
			return errors.ErrorFileTooLarge
		}

		fileType, err := fileops.DetectMultipartFileType(req.Image3)

		if err != nil {
			return errors.ErrorInvalidFileType
		}

		allowedTypes := fileops.ImageContentTypes
		if !slices.Contains(allowedTypes, fileType) {
			return errors.ErrorInvalidFileType
		}

		imageURL3, err = ps.supabase.Upload(req.Image3)
		if err != nil {
			return err
		}
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
		ImageURL1: imageURL1,
		ImageURL2: imageURL2,
		ImageURL3: imageURL3,
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

	place := entity.Place{ID: req.ID}
	rowsAffected, err := ps.pr.GetByID(&place)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.ErrorNotFound
	}

	var imageURL1 = ""
	var imageURL2 = ""
	var imageURL3 = ""

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
		
		ps.supabase.Delete(place.ImageURL1)

		imageURL1, err = ps.supabase.Upload(req.Image1)
		if err != nil {
			return err
		}
	}
	if req.Image2 != nil {
		if req.Image2.Size > 1*fileops.MegaByte {
			return errors.ErrorFileTooLarge
		}

		fileType, err := fileops.DetectMultipartFileType(req.Image2)

		if err != nil {
			return errors.ErrorInvalidFileType
		}

		allowedTypes := fileops.ImageContentTypes
		if !slices.Contains(allowedTypes, fileType) {
			return errors.ErrorInvalidFileType
		}

		ps.supabase.Delete(place.ImageURL2)

		imageURL2, err = ps.supabase.Upload(req.Image2)
		if err != nil {
			return err
		}
	}
	if req.Image3 != nil {
		if req.Image3.Size > 1*fileops.MegaByte {
			return errors.ErrorFileTooLarge
		}

		fileType, err := fileops.DetectMultipartFileType(req.Image3)

		if err != nil {
			return errors.ErrorInvalidFileType
		}

		allowedTypes := fileops.ImageContentTypes
		if !slices.Contains(allowedTypes, fileType) {
			return errors.ErrorInvalidFileType
		}

		ps.supabase.Delete(place.ImageURL3)

		imageURL3, err = ps.supabase.Upload(req.Image3)
		if err != nil {
			return err
		}
	}

	place = entity.Place{
		ID: req.ID,
		Name: req.Name,
		Location: req.Location,
		Description: req.Description,
		Address: req.Address,
		OpeningHours: req.OpeningHours,
		ClosingHours: req.ClosingHours,
		EntryPrice: req.EntryPrice,
		ImageURL1: imageURL1,
		ImageURL2: imageURL2,
		ImageURL3: imageURL3,
		MapURL: req.MapURL,
		Rating: req.Rating,
	}

	rowsAffected, err = ps.pr.Update(&place)
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