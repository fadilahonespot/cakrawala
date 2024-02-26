package entity

import (
	"time"

	"gorm.io/gorm"
)

type City struct {
	ID              int `gorm:"primarykey"`
	Name            string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	User            []User         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OriginCity      []Courier      `gorm:"foreignKey:OriginCityID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DestinationCity []Courier      `gorm:"foreignKey:DestinationCityID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
