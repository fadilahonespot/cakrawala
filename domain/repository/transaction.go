package repository

import (
	"context"

	"github.com/fadilahonespot/cakrawala/domain/entity"
	"github.com/fadilahonespot/cakrawala/utils/paginate"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	BeginTrans(ctx context.Context) *gorm.DB
	CreateTransactionBasket(ctx context.Context, req *entity.TransactionBasket) (err error)
	UpdateTransactionBasket(ctx context.Context, tx *gorm.DB, req *entity.TransactionBasket) (err error)
	FindTransactionBasketByUserIdStatusPending(ctx context.Context, userId string) (resp *entity.TransactionBasket, err error)
	FindAllTransactionBasketByUserId(ctx context.Context, userId string, params paginate.Pagination) (resp []entity.TransactionBasket, count int64, err error)
	FindTransactionDetailByBasketId(ctx context.Context, basketId string) (resp []entity.TransactionDetail, err error)
	CreateTransactionDetail(ctx context.Context, tx *gorm.DB, req *entity.TransactionDetail) (err error)
	DeleteTransactionDetail(ctx context.Context, tx *gorm.DB, id int) (err error)
	FindTransactionDetailByProductId(ctx context.Context, productId string) (resp entity.TransactionDetail, err error)
	UpdateTransactionDetail(ctx context.Context, tx *gorm.DB, req *entity.TransactionDetail) (err error)
	CreateTransaction(ctx context.Context, tx *gorm.DB, req *entity.Transaction) (err error)
	FindTransactionDetailByProductIdAndBasketId(ctx context.Context, productId, basketId string) (resp entity.TransactionDetail, err error)
	FindTransactionById(ctx context.Context, id string) (resp *entity.Transaction, err error) 
}
