package repository

import (
	"context"

	"github.com/fadilahonespot/cakrawala/domain/entity"
)

type CourierInfoRepository interface {
	FindByCode(ctx context.Context, code string) (resp *entity.CourierInfo, err error)
	FindAll(ctx context.Context) (resp []entity.CourierInfo, err error)
}