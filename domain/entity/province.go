package entity

import (
	"time"

	"gorm.io/gorm"
)

type Province struct {
	ID        int `gorm:"primarykey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	User      []User         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
