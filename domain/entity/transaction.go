package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID                  string `gorm:"primaryKey"`
	TransactionBasketID int
	TransactionBasket   TransactionBasket `gorm:"foreignKey:TransactionBasketID"`
	TotalQty            int
	TotalPrice          float64
	TotalWeight         int
	PaymentInfoID       int
	PaymentInfo         PaymentInfo `gorm:"foreignKey:PaymentInfoID"`
	CourierID           int
	Courier             Courier `gorm:"foreignKey:CourierID"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}

func (m *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.New().String()
	m.ID = uuid

	return nil
}
