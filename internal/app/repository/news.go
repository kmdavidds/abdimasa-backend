package repository

import (
	"github.com/google/uuid"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/entity"
	"gorm.io/gorm"
)

type NewsRepository interface {
	Create(news *entity.News) (int64, error)
	GetAll(news *[]entity.News) (int64, error)
	GetByID(news *entity.News) (int64, error)
	Update(news *entity.News) (int64, error)
	Delete(id uuid.UUID) (int64, error)
}

type newsRepository struct {
	db *gorm.DB
}

func NewNewsRepository(db *gorm.DB) NewsRepository {
	return &newsRepository{db}
}

func (nr *newsRepository) Create(news *entity.News) (int64, error) {
	result := nr.db.Debug().Create(news)
	return result.RowsAffected, result.Error
}

func (nr *newsRepository) GetAll(news *[]entity.News) (int64, error) {
	result := nr.db.Debug().Find(news)
	return result.RowsAffected, result.Error
}

func (nr *newsRepository) GetByID(news *entity.News) (int64, error) {
	result := nr.db.Debug().First(news)
	return result.RowsAffected, result.Error
}

func (nr *newsRepository) Update(news *entity.News) (int64, error) {
	result := nr.db.Debug().Model(&entity.News{}).Where("id = ?", news.ID).Updates(news)
	return result.RowsAffected, result.Error
}

func (nr *newsRepository) Delete(id uuid.UUID) (int64, error) {
	result := nr.db.Debug().Where("id = ?", id).Delete(&entity.News{})
	return result.RowsAffected, result.Error
}

