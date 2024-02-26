package repository

import (
	"context"

	"github.com/fadilahonespot/cakrawala/domain/entity"
)

type ProductTypeRepository interface {
	CreateProductType(ctx context.Context, productType *entity.ProductType) (err error)
	UpdateProductType(ctx context.Context, productType *entity.ProductType) (err error)
	DeleteProductType(ctx context.Context, id string) (err error)
	GetProductType(ctx context.Context, id string) (productType *entity.ProductType, err error)
	GetAllProductTypes(ctx context.Context) (resp []entity.ProductType, err error)
}