package repository

import (
	"github.com/google/uuid"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/entity"
	"gorm.io/gorm"
)

type PlaceRepository interface {
	Create(place *entity.Place) (int64, error)
	GetAll(places *[]entity.Place) (int64, error)
	GetByID(place *entity.Place) (int64, error)
	Update(place *entity.Place) (int64, error)
	Delete(id uuid.UUID) (int64, error)
}

type placeRepository struct {
	db *gorm.DB
}

func NewPlaceRepository(db *gorm.DB) PlaceRepository {
	return &placeRepository{db}
}

func (pr *placeRepository) Create(place *entity.Place) (int64, error) {
	result := pr.db.Debug().Create(place)
	return result.RowsAffected, result.Error
}

func (pr *placeRepository) GetAll(places *[]entity.Place) (int64, error) {
	result := pr.db.Debug().Find(places)
	return result.RowsAffected, result.Error
}

func (pr *placeRepository) GetByID(place *entity.Place) (int64, error) {
	result := pr.db.Debug().First(place)
	return result.RowsAffected, result.Error
}

func (pr *placeRepository) Update(place *entity.Place) (int64, error) {
	result := pr.db.Debug().Model(&entity.Place{}).Where("id = ?", place.ID).Updates(place)
	return result.RowsAffected, result.Error
}

func (pr *placeRepository) Delete(id uuid.UUID) (int64, error) {
	result := pr.db.Debug().Where("id = ?", id).Delete(&entity.Place{})
	return result.RowsAffected, result.Error
}