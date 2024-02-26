package entity

import (
	"time"

	"gorm.io/gorm"
)

type Courier struct {
	ID                int `gorm:"primarykey"`
	OriginCityID      int
	OriginCity        City `gorm:"foreignKey:OriginCityID"`
	DestinationCityID int
	DestinationCity   City `gorm:"foreignKey:DestinationCityID"`
	CourierInfoID     int
	CourierInfo       CourierInfo `gorm:"foreignKey:CourierInfoID"`
	Service           string
	Description       string
	Price             int
	ETD               string
	Note              string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	Transaction       []Transaction  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
