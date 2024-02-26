package entity

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID                int `gorm:"primarykey"`
	Name              string
	Code              string
	IsAvailable       bool
	Description       string
	ProductTypeID     int
	ProductType       ProductType `gorm:"foreignKey:ProductTypeID"`
	Weight            int
	Price             float64
	Stock             int
	Sold              int
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt      `gorm:"index"`
	ProductImg        []ProductImg        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TransactionDetail []TransactionDetail `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
