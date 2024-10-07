package repository

import (
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/entity"
	"gorm.io/gorm"
)

type DetailRepository interface {
	Create(detail *entity.Detail) (int64, error)
	GetAll(detail *[]entity.Detail) (int64, error)
	GetByID(detail *entity.Detail) (int64, error)
	GetBySlug(detail *entity.Detail) (int64, error)
	Update(detail *entity.Detail) (int64, error)
	Delete(id uint8) (int64, error)
}

type detailRepository struct {
	db *gorm.DB
}

func NewDetailRepository(db *gorm.DB) DetailRepository {
	return &detailRepository{db}
}

func (dr *detailRepository) Create(detail *entity.Detail) (int64, error) {
	result := dr.db.Debug().Create(detail)
	return result.RowsAffected, result.Error
}

func (dr *detailRepository) GetAll(detail *[]entity.Detail) (int64, error) {
	result := dr.db.Debug().Find(detail)
	return result.RowsAffected, result.Error
}

func (dr *detailRepository) GetByID(detail *entity.Detail) (int64, error) {
	result := dr.db.Debug().First(detail)
	return result.RowsAffected, result.Error
}

func (dr *detailRepository) GetBySlug(detail *entity.Detail) (int64, error) {
	result := dr.db.Debug().First(detail)
	return result.RowsAffected, result.Error
}

func (dr *detailRepository) Update(detail *entity.Detail) (int64, error) {
	result := dr.db.Debug().Model(&entity.Detail{}).Where("id = ?", detail.ID).Updates(detail)
	return result.RowsAffected, result.Error
}

func (dr *detailRepository) Delete(id uint8) (int64, error) {
	result := dr.db.Debug().Where("id = ?", id).Delete(&entity.Detail{})
	return result.RowsAffected, result.Error
}

