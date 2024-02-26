package repository

import (
	"context"

	"github.com/fadilahonespot/cakrawala/domain/entity"
	"gorm.io/gorm"
)

type CourierRepository interface {
	Create(ctx context.Context, tx *gorm.DB, req *entity.Courier) (err error)
}