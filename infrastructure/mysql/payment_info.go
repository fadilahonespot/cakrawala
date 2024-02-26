package mysql

import (
	"context"

	"github.com/fadilahonespot/cakrawala/domain/entity"
	"github.com/fadilahonespot/cakrawala/domain/repository"
	"gorm.io/gorm"
)

type defaultPaymentInfo struct {
	db *gorm.DB
}

func SetupPaymentInfoRepository(db *gorm.DB) repository.PaymentInfoRepository {
	return &defaultPaymentInfo{db: db}
}

func (r *defaultPaymentInfo) Create(ctx context.Context, req *entity.PaymentInfo) (err error) {
	err = r.db.WithContext(ctx).Create(req).Error
	return
}

func (r *defaultPaymentInfo) FindByXpayment(ctx context.Context, externalId string) (resp *entity.PaymentInfo, err error) {
	err = r.db.WithContext(ctx).Take(&resp, "x_payment = ?", externalId).Error
	return
}

func (r *defaultPaymentInfo) Update(ctx context.Context, req *entity.PaymentInfo) (err error) {
	err = r.db.WithContext(ctx).Save(req).Error
	return
}