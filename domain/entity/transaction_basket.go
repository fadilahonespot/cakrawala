package entity

import (
	"time"

	"gorm.io/gorm"
)

type BasketStatus string

const (
	BasketPending   BasketStatus = "Pending"
	BasketCompleted BasketStatus = "Completed"
)

type TransactionBasket struct {
	ID                int `gorm:"primarykey"`
	UserID            int
	User              User `gorm:"foreignKey:UserID"`
	BasketStatus      BasketStatus
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	TransactionDetail []TransactionDetail
	Transaction       []Transaction
}
