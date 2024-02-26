package repository

import (
	"context"

	"github.com/fadilahonespot/cakrawala/domain/entity"
)

type PaymentInfoRepository interface {
	Create(ctx context.Context, req *entity.PaymentInfo) (err error)
	FindByXpayment(ctx context.Context, externalId string) (resp *entity.PaymentInfo, err error)
	Update(ctx context.Context, req *entity.PaymentInfo) (err error)
}