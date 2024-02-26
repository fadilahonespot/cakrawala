package database

import (
	"github.com/fadilahonespot/cakrawala/domain/entity"
	"gorm.io/gorm"
)

func initMigrate(db *gorm.DB) {
	db.AutoMigrate(&entity.Province{})
    db.AutoMigrate(&entity.City{})
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.ProductType{})
	db.AutoMigrate(&entity.Product{})
	db.AutoMigrate(&entity.ProductImg{})
	db.AutoMigrate(&entity.CourierInfo{})
	db.AutoMigrate(&entity.Courier{})
	db.AutoMigrate(&entity.PaymentInfo{})
	db.AutoMigrate(&entity.TransactionBasket{})
	db.AutoMigrate(&entity.Transaction{})
	db.AutoMigrate(&entity.TransactionDetail{})
}