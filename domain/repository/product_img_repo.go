package repository

import (
	"context"

	"github.com/fadilahonespot/cakrawala/domain/entity"
	"gorm.io/gorm"
)

type ProductImgRepository interface {
	UpdateProductImg(ctx context.Context, tx *gorm.DB, req *entity.ProductImg) (err error)
	DeleteProductImg(ctx context.Context, tx *gorm.DB, id string) (err error)
	GetProductImgByProductId(ctx context.Context, id string) (resp []entity.ProductImg, err error)
	DeleteProductImgByProductId(ctx context.Context, tx *gorm.DB, id string) (err error)
}