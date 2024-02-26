package mysql

import (
	"context"

	"github.com/fadilahonespot/cakrawala/domain/entity"
	"github.com/fadilahonespot/cakrawala/domain/repository"
	"gorm.io/gorm"
)

type defaultCourierInfoRepo struct {
	db *gorm.DB
}

func SetupCourierInfoRepository(db *gorm.DB) repository.CourierInfoRepository {
	return &defaultCourierInfoRepo{db: db}
}

func (r *defaultCourierInfoRepo) FindByCode(ctx context.Context, code string) (resp *entity.CourierInfo, err error) {
	err = r.db.WithContext(ctx).Take(&resp, "code = ?", code).Error
	return
}

func (r *defaultCourierInfoRepo) FindAll(ctx context.Context) (resp []entity.CourierInfo, err error) {
	err = r.db.WithContext(ctx).Find(&resp).Error
	return
}
 

