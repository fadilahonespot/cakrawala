package entity

import (
	"time"

	"gorm.io/gorm"
)

type CourierInfo struct {
	ID        int `gorm:"primarykey"`
	Name      string
	Code      string
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Courier   []Courier      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
