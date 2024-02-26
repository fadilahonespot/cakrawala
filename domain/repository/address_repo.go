package repository

import (
	"context"

	"github.com/fadilahonespot/cakrawala/domain/entity"
)

type AddressRepository interface {
	GetProvince(ctx context.Context, id string) (*entity.Province, error)
	UpdateProvince(ctx context.Context, req *entity.Province) error 
	GetCity(ctx context.Context, id string) (*entity.City, error) 
	UpdateCity(ctx context.Context, req *entity.City) error
}