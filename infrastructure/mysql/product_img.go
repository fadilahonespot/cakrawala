package mysql

import (
	"context"

	"github.com/fadilahonespot/cakrawala/domain/entity"
	"github.com/fadilahonespot/cakrawala/domain/repository"
	"gorm.io/gorm"
)

type defaultProductImg struct {
	db *gorm.DB
}

func SetupProductImgRepository(db *gorm.DB) repository.ProductImgRepository {
	return &defaultProductImg{
		db: db,
	}
}

func (r *defaultProductImg) UpdateProductImg(ctx context.Context, tx *gorm.DB, req *entity.ProductImg) (err error) {
	err = tx.WithContext(ctx).Save(req).Error
	return
}

func (r *defaultProductImg) DeleteProductImg(ctx context.Context, tx *gorm.DB, id string) (err error) {
	err = r.db.WithContext(ctx).Delete(&entity.ProductImg{}, "id =?", id).Error
	return
}

func (r *defaultProductImg) GetProductImgByProductId(ctx context.Context, id string) (resp []entity.ProductImg, err error) {
	err = r.db.WithContext(ctx).Take(&resp, "product_id = ?", id).Error
	return
}

func (r *defaultProductImg) DeleteProductImgByProductId(ctx context.Context, tx *gorm.DB, id string) (err error) {
	err = tx.WithContext(ctx).Delete(&entity.ProductImg{}, "product_id =?", id).Error
	return
}

