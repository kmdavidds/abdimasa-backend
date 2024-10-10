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

type BusinessService interface {
	Create(req dto.CreateBusinessRequest) error
	GetAll() ([]entity.Business, error)
	GetByID(req dto.GetBusinessByIDRequest) (entity.Business, error)
	Update(req dto.UpdateBusinessRequest) error
	Delete(req dto.DeleteBusinessRequest) error
}

type businessService struct {
	br       repository.BusinessRepository
	val      validator.Validator
	supabase supabase.Supabase
}

func NewBusinessService(
	br repository.BusinessRepository,
	val validator.Validator,
	supabase supabase.Supabase,
) BusinessService {
	return &businessService{br, val, supabase}
}

func (bs *businessService) Create(req dto.CreateBusinessRequest) error {
	valErr := bs.val.Validate(req)
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

		imageURL1, err = bs.supabase.Upload(req.Image1)
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

		imageURL2, err = bs.supabase.Upload(req.Image2)
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

		imageURL3, err = bs.supabase.Upload(req.Image3)
		if err != nil {
			return err
		}
	}

	idV7, err := uuid.NewV7()
	if err != nil {
		return err
	}

	business := entity.Business{
		ID:          idV7,
		Name:        req.Name,
		Location:    req.Location,
		Description: req.Description,
		Address:     req.Address,
		PriceRange:  req.PriceRange,
		ImageURL1:   imageURL1,
		ImageURL2:   imageURL2,
		ImageURL3:   imageURL3,
		Contact:     req.Contact,
		MapURL:      req.MapURL,
		Rating:      req.Rating,
	}

	_, err = bs.br.Create(&business)
	if err != nil {
		return err
	}

	return nil
}

func (bs *businessService) GetAll() ([]entity.Business, error) {
	businesses := []entity.Business{}
	_, err := bs.br.GetAll(&businesses)
	if err != nil {
		return nil, err
	}

	return businesses, nil
}

func (bs *businessService) GetByID(req dto.GetBusinessByIDRequest) (entity.Business, error) {
	valErr := bs.val.Validate(req)
	if valErr != nil {
		return entity.Business{}, valErr
	}

	business := entity.Business{ID: req.ID}

	rowsAffected, err := bs.br.GetByID(&business)
	if err != nil {
		return entity.Business{}, err
	}

	if rowsAffected == 0 {
		return entity.Business{}, errors.ErrorNotFound
	}

	return business, nil
}

func (bs *businessService) Update(req dto.UpdateBusinessRequest) error {
	valErr := bs.val.Validate(req)
	if valErr != nil {
		return valErr
	}

	business := entity.Business{ID: req.ID}
	rowsAffected, err := bs.br.GetByID(&business)
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
		
		bs.supabase.Delete(business.ImageURL1)

		imageURL1, err = bs.supabase.Upload(req.Image1)
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

		bs.supabase.Delete(business.ImageURL2)

		imageURL2, err = bs.supabase.Upload(req.Image2)
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

		bs.supabase.Delete(business.ImageURL3)

		imageURL3, err = bs.supabase.Upload(req.Image3)
		if err != nil {
			return err
		}
	}

	business = entity.Business{
		ID:          req.ID,
		Name:        req.Name,
		Location:    req.Location,
		Description: req.Description,
		Address:     req.Address,
		PriceRange:  req.PriceRange,
		ImageURL1:   imageURL1,
		ImageURL2:   imageURL2,
		ImageURL3:   imageURL3,
		Contact:     req.Contact,
		MapURL:      req.MapURL,
		Rating:      req.Rating,
	}

	rowsAffected, err = bs.br.Update(&business)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.ErrorNotFound
	}

	return nil
}

func (bs *businessService) Delete(req dto.DeleteBusinessRequest) error {
	valErr := bs.val.Validate(req)
	if valErr != nil {
		return valErr
	}

	rowsAffected, err := bs.br.Delete(req.ID)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.ErrorNotFound
	}

	return nil
}
