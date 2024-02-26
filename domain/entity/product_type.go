package entity

import (
	"time"

	"gorm.io/gorm"
)

type ProductType struct {
	ID        int `gorm:"primarykey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Product   []Product      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
