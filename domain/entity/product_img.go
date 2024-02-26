package entity

import (
	"time"

	"gorm.io/gorm"
)

type ProductImg struct {
	ID        int `gorm:"primarykey"`
	Image     string
	ProductID int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
