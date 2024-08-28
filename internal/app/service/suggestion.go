package service

import (
	"github.com/google/uuid"
	"github.com/kmdavidds/abdimasa-backend/internal/app/repository"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/dto"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/entity"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/errors"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/validator"
)

type SuggestionService interface {
	Create(req dto.CreateSuggestionRequest) error
	GetAll() ([]entity.Suggestion, error)
	Delete(req dto.DeleteSuggestionRequest) error
}

type suggestionService struct {
	sr  repository.SuggestionRepository
	val validator.Validator
}

func NewSuggestionService(
	sr repository.SuggestionRepository,
	val validator.Validator,
) SuggestionService {
	return &suggestionService{sr, val}
}

func (ss *suggestionService) Create(req dto.CreateSuggestionRequest) error {
	valErr := ss.val.Validate(req)
	if valErr != nil {
		return valErr
	}

	idV7, err := uuid.NewV7()
	if err != nil {
		return err
	}

	suggestion := entity.Suggestion{
		ID:            idV7,
		Name:          req.Name,
		Description:   req.Description,
		AttachmentURL: req.AttachmentURL,
	}

	_, err = ss.sr.Create(&suggestion)
	if err != nil {
		return err
	}

	return nil
}

func (ss *suggestionService) GetAll() ([]entity.Suggestion, error) {
	suggestions := []entity.Suggestion{}
	_, err := ss.sr.GetAll(&suggestions)
	if err != nil {
		return nil, err
	}

	return suggestions, nil
}

func (ss *suggestionService) Delete(req dto.DeleteSuggestionRequest) error {
	valErr := ss.val.Validate(req)
	if valErr != nil {
		return valErr
	}

	rowsAffected, err := ss.sr.Delete(req.ID)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.ErrorNotFound
	}

	return nil
}
