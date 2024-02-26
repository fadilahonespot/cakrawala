package mysql

import (
	"context"

	"github.com/fadilahonespot/cakrawala/domain/entity"
	"github.com/fadilahonespot/cakrawala/domain/repository"
	"gorm.io/gorm"
)

type defaultCourierRepo struct {
	db *gorm.DB
}

func SetupCourierRepository(db *gorm.DB) repository.CourierRepository {
	return &defaultCourierRepo{db: db}
}

func (r *defaultCourierRepo) Create(ctx context.Context, tx *gorm.DB, req *entity.Courier) (err error) {
	err = r.db.WithContext(ctx).Create(req).Error
	return
}
