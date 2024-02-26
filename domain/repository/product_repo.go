package repository

import (
	"context"

	"github.com/fadilahonespot/cakrawala/domain/entity"
	"github.com/fadilahonespot/cakrawala/utils/paginate"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(ctx context.Context, tx *gorm.DB, product *entity.Product) (err error)
	FindAll(ctx context.Context, param paginate.Pagination) (products []entity.Product, count int64, err error)
	FindByID(ctx context.Context, id int) (product *entity.Product, err error)
	Update(ctx context.Context, tx *gorm.DB, product *entity.Product) (err error)
	Delete(ctx context.Context, tx *gorm.DB, id int) (err error)
	BeginTrans(ctx context.Context) *gorm.DB
	FindProductByCode(ctx context.Context, code string) (resp *entity.Product, err error)
}