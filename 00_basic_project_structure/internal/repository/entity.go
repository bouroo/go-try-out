package repository

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Task struct {
	ID          uint           `gorm:"primaryKey"`
	Title       string         `gorm:"index;size:127"`
	Description datatypes.JSON `gorm:"default:'{}'"`
	Status      string         `gorm:"default:'pending'"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
