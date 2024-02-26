package mysql

import (
	"context"

	"github.com/fadilahonespot/cakrawala/domain/entity"
	"github.com/fadilahonespot/cakrawala/domain/repository"
	"gorm.io/gorm"
)

type defaultProductType struct {
	db *gorm.DB
}

func SetupProductTypeRepository(db *gorm.DB) repository.ProductTypeRepository {
	return &defaultProductType{
        db: db,
    }
}

func (r *defaultProductType) GetProductType(ctx context.Context, id string) (resp *entity.ProductType, err error) {
	err = r.db.WithContext(ctx).First(&resp, "id =?", id).Error
    return
}

func (r *defaultProductType) UpdateProductType(ctx context.Context, req *entity.ProductType) (err error) {
    err = r.db.WithContext(ctx).Save(req).Error
	return
}

func (r *defaultProductType) DeleteProductType(ctx context.Context, id string) (err error) {
	err = r.db.WithContext(ctx).Delete(&entity.ProductType{}, "id =?", id).Error
    return
}

func (r *defaultProductType) CreateProductType(ctx context.Context, req *entity.ProductType) (err error) {
	err = r.db.WithContext(ctx).Create(req).Error
    return
}

func (r *defaultProductType) GetAllProductTypes(ctx context.Context) (resp []entity.ProductType, err error) {
	err = r.db.WithContext(ctx).Find(&resp).Error
    return
}  