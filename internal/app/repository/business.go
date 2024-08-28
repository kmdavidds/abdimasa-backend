package repository

import (
	"github.com/google/uuid"

	"github.com/kmdavidds/abdimasa-backend/internal/pkg/entity"
	"gorm.io/gorm"
)

type BusinessRepository interface {
	Create(business *entity.Business) (int64, error)
	GetAll(businesses *[]entity.Business) (int64, error)
	Update(business *entity.Business) (int64, error)
	Delete(id uuid.UUID) (int64, error)
}

type businessRepository struct {
	db *gorm.DB
}

func NewBusinessRepository(db *gorm.DB) BusinessRepository {
	return &businessRepository{db}
}

func (br *businessRepository) Create(business *entity.Business) (int64, error) {
	result := br.db.Debug().Create(business)
	return result.RowsAffected, result.Error
}

func (br *businessRepository) GetAll(businesses *[]entity.Business) (int64, error) {
	result := br.db.Debug().Find(businesses)
	return result.RowsAffected, result.Error
}

func (br *businessRepository) Update(business *entity.Business) (int64, error) {
	result := br.db.Debug().Model(&entity.Business{}).Where("id = ?", business.ID).Updates(business)
	return result.RowsAffected, result.Error
}

func (br *businessRepository) Delete(id uuid.UUID) (int64, error) {
	result := br.db.Debug().Where("id = ?", id).Delete(&entity.Business{})
	return result.RowsAffected, result.Error
}