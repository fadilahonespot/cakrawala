package mysql

import (
	"context"

	"github.com/fadilahonespot/cakrawala/domain/entity"
	"github.com/fadilahonespot/cakrawala/domain/repository"
	"gorm.io/gorm"
)

type defaultAddressRepo struct {
	db *gorm.DB
}

func SetupAddressRepoRepository(db *gorm.DB) repository.AddressRepository {
	return &defaultAddressRepo{
        db: db,
    }
}

func (r *defaultAddressRepo) GetProvince(ctx context.Context, id string) (*entity.Province, error) {
	var resp entity.Province
	err := r.db.WithContext(ctx).First(&resp, "id = ?", id).Error
	return &resp, err
}

func (r *defaultAddressRepo) UpdateProvince(ctx context.Context, req *entity.Province) error {
    return r.db.WithContext(ctx).Save(req).Error
}

func (r *defaultAddressRepo) GetCity(ctx context.Context, id string) (*entity.City, error) {
	var resp entity.City
	err := r.db.WithContext(ctx).First(&resp, "id = ?", id).Error
	return &resp, err
}

func (r *defaultAddressRepo) UpdateCity(ctx context.Context, req *entity.City) error {
    return r.db.WithContext(ctx).Save(req).Error
}
