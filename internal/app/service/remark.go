package service

import (
	"github.com/google/uuid"
	"github.com/kmdavidds/abdimasa-backend/internal/app/repository"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/dto"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/entity"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/validator"
)

type RemarkService interface {
	Create(req dto.CreateRemarkRequest) error
	GetAll() ([]entity.Remark, error)
}

type remarkService struct {
	rr  repository.RemarkRepository
	val validator.Validator
}

func NewRemarkService(
	rr repository.RemarkRepository,
	val validator.Validator,
) RemarkService {
	return &remarkService{rr, val}
}

func (rs *remarkService) Create(req dto.CreateRemarkRequest) error {
	valErr := rs.val.Validate(req)
	if valErr != nil {
		return valErr
	}

	idV7, err := uuid.NewV7()
	if err != nil {
		return err
	}

	remark := entity.Remark{
		ID:         idV7,
		Name:       req.Name,
		Occupation: req.Occupation,
		Description: req.Description,
		Gender:     req.Gender,
	}

	_, err = rs.rr.Create(&remark)
	if err != nil {
		return err
	}

	return nil
}

func (rs *remarkService) GetAll() ([]entity.Remark, error) {
	remarks := []entity.Remark{}
	_, err := rs.rr.GetAll(&remarks)
	if err != nil {
		return nil, err
	}

	return remarks, nil
}
