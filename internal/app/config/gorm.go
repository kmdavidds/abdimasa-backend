package config

import (
	"log"
	"os"

	"github.com/kmdavidds/abdimasa-backend/internal/pkg/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("POSTGRES_DSN")), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatalf("failed to connect to database %v", err)
	}

	return db
}

func MigrateTables(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entity.Activity{},
		&entity.Place{},
		&entity.Business{},
		&entity.Remark{},
	)
	if err != nil {
		return err
	}

	return nil
}