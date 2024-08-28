package repository

import (
	"github.com/google/uuid"

	"github.com/kmdavidds/abdimasa-backend/internal/pkg/entity"
	"gorm.io/gorm"
)

type SuggestionRepository interface {
	Create(suggestion *entity.Suggestion) (int64, error)
	GetAll(suggestions *[]entity.Suggestion) (int64, error)
	Delete(id uuid.UUID) (int64, error)
}

type suggestionRepository struct {
	db *gorm.DB
}

func NewSuggestionRepository(db *gorm.DB) SuggestionRepository {
	return &suggestionRepository{db}
}

func (sr *suggestionRepository) Create(suggestion *entity.Suggestion) (int64, error) {
	result := sr.db.Debug().Create(suggestion)
	return result.RowsAffected, result.Error
}

func (sr *suggestionRepository) GetAll(suggestions *[]entity.Suggestion) (int64, error) {
	result := sr.db.Debug().Find(suggestions)
	return result.RowsAffected, result.Error
}

func (sr *suggestionRepository) Delete(id uuid.UUID) (int64, error) {
	result := sr.db.Debug().Where("id = ?", id).Delete(&entity.Suggestion{})
	return result.RowsAffected, result.Error
}