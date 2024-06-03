package repository

import (
	"log/slog"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	err := autoMigrate(db)
	if err != nil {
		slog.Error("NewRepository", "AutoMigrate", err)
	}
	return &Repository{
		db: db,
	}
}

func autoMigrate(db *gorm.DB) (err error) {
	err = db.AutoMigrate(&Task{})
	return
}
