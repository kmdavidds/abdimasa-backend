package repository

import (
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/entity"
	"gorm.io/gorm"
)

type RemarkRepository interface {
	Create(remark *entity.Remark) (int64, error)
	GetAll(remarks *[]entity.Remark) (int64, error)
}

type remarkRepository struct {
	db *gorm.DB
}

func NewRemarkRepository(db *gorm.DB) RemarkRepository {
	return &remarkRepository{db}
}

func (rr *remarkRepository) Create(remark *entity.Remark) (int64, error) {
	result := rr.db.Debug().Create(remark)
	return result.RowsAffected, result.Error
}

func (rr *remarkRepository) GetAll(remarks *[]entity.Remark) (int64, error) {
	result := rr.db.Debug().Find(remarks)
	return result.RowsAffected, result.Error
}