package entity

import (
	"time"

	"gorm.io/gorm"
)

type TransactionDetail struct {
	ID                  int `gorm:"primarykey"`
	TransactionBasketID int
	TransactionBasket   TransactionBasket `gorm:"foreignKey:TransactionBasketID"`
	ProductID           int
	Product             Product `gorm:"foreignKey:ProductID"`
	Qty                 int
	Price               float64
	Weight              int
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}
