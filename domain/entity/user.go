package entity

import (
	"time"

	"gorm.io/gorm"
)

type EmailStatus string

const (
	Unverify EmailStatus = "UNVERIFY"
	Verify   EmailStatus = "VERIFY"
)

type RoleName string

const (
	Admin      RoleName = "admin"
	Customer   RoleName = "customer"
)

type User struct {
	ID                int `gorm:"primarykey"`
	Name              string
	PhoneNumber       string
	Email             string
	EmailStatus       EmailStatus
	Password          string
	Address           string
	CityID            int
	ProvinceID        int
	PostalCode        string
	Role              RoleName
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt      `gorm:"index"`
	TransactionBasket []TransactionBasket `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
