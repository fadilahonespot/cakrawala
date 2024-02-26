package entity

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type PaymentStatus string

const (
	Failed PaymentStatus = "failed"
	Unpaid PaymentStatus = "unpaid"
	Paid   PaymentStatus = "paid"
)

type PaymentInfo struct {
	ID             int `gorm:"primarykey"`
	XPayment       string
	Status         PaymentStatus
	BankCode       string
	AccountNumber  string
	Amount         float64
	PaymentTime    *time.Time
	ExpirationDate time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	Transaction    []Transaction  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (p *PaymentInfo) BeforeSave(tx *gorm.DB) (err error) {
	if p.Status == Paid {
		err = tx.Take(&p.Transaction, "payment_info_id = ?", p.ID).Error
		if err != nil {
			return errors.New("[before save] failed to find transaction by payment info")
		}
		var transactionBasket TransactionBasket
		if len(p.Transaction) != 0 {
			err := tx.Preload("TransactionDetail").Preload("TransactionDetail.Product").
				Take(&transactionBasket, "id = ?", p.Transaction[0].TransactionBasketID).Error
			if err != nil {
				return errors.New("[before save] error find transaction basket")
			}
			for _, detail := range transactionBasket.TransactionDetail {
				product := detail.Product

				product.Stock -= detail.Qty
				product.Sold += detail.Qty

				if err := tx.Save(&product).Error; err != nil {
					return err
				}
			}

		}
	}

	return
}
