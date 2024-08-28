package repository

import (
	"github.com/google/uuid"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/entity"
	"gorm.io/gorm"
)

type ActivityRepository interface {
	Create(activity *entity.Activity) (int64, error)
	GetAll(activities *[]entity.Activity) (int64, error)
	Update(activity *entity.Activity) (int64, error)
	Delete(id uuid.UUID) (int64, error)
}

type activityRepository struct {
	db *gorm.DB
}

func NewActivityRepository(db *gorm.DB) ActivityRepository {
	return &activityRepository{db}
}

func (ar *activityRepository) Create(activity *entity.Activity) (int64, error) {
	result := ar.db.Debug().Create(activity)
	return result.RowsAffected, result.Error
}

func (ar *activityRepository) GetAll(activities *[]entity.Activity) (int64, error) {
	result := ar.db.Debug().Find(activities)
	return result.RowsAffected, result.Error
}

func (ar *activityRepository) Update(activity *entity.Activity) (int64, error) {
	result := ar.db.Debug().Model(&entity.Activity{}).Where("id = ?", activity.ID).Updates(activity)
	return result.RowsAffected, result.Error
}

func (ar *activityRepository) Delete(id uuid.UUID) (int64, error) {
	result := ar.db.Debug().Where("id = ?", id).Delete(&entity.Activity{})
	return result.RowsAffected, result.Error
}