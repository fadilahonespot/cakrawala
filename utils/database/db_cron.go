package database

import (
	"fmt"
	"time"

	"github.com/fadilahonespot/cakrawala/domain/entity"
	"gorm.io/gorm"
)

func tiker(db *gorm.DB) {
	fmt.Println("Starting Cron Job Database....")
	ticker := time.NewTicker(5 * time.Hour) // Ubah sesuai dengan interval yang Anda inginkan
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            updatePaymentStatus(db)
        }
    }
}

func updatePaymentStatus(db *gorm.DB) {
	currentTime := time.Now()

	db.Model(&entity.PaymentInfo{}).
		Where("status = ?", entity.Unpaid).
		Where("payment_time IS NULL").
		Where("expiration_date <= ?", currentTime).
		Update("status", entity.Failed)
}