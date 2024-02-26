package mysql

import (
	"context"

	"github.com/fadilahonespot/cakrawala/domain/entity"
	"github.com/fadilahonespot/cakrawala/domain/repository"
	"github.com/fadilahonespot/cakrawala/utils/paginate"
	"gorm.io/gorm"
)

type defaultTransaction struct {
	db *gorm.DB
}

func SetupTransactionRepository(db *gorm.DB) repository.TransactionRepository {
	return &defaultTransaction{db: db}
}

func (s *defaultTransaction) BeginTrans(ctx context.Context) *gorm.DB {
	return s.db.Begin()
}

func (s *defaultTransaction) CreateTransactionBasket(ctx context.Context, req *entity.TransactionBasket) (err error) {
	err = s.db.WithContext(ctx).Create(req).Error
	return
}

func (s *defaultTransaction) UpdateTransactionBasket(ctx context.Context, tx *gorm.DB, req *entity.TransactionBasket) (err error) {
	err = tx.WithContext(ctx).Save(req).Error
	return
}

func (s *defaultTransaction) FindTransactionBasketByUserIdStatusPending(ctx context.Context, userId string) (resp *entity.TransactionBasket, err error) {
	err = s.db.WithContext(ctx).Take(&resp, "user_id = ? AND basket_status = ?", userId, entity.BasketPending).Error
	return
}

func (s *defaultTransaction) FindAllTransactionBasketByUserId(ctx context.Context, userId string, params paginate.Pagination) (resp []entity.TransactionBasket, count int64, err error) {
	err = s.db.WithContext(ctx).Model(&entity.TransactionBasket{}).Count(&count).Error
	if err != nil {
		return
	}
	err = s.db.WithContext(ctx).Order("id DESC").Preload("Transaction").Preload("Transaction.Courier").
		Preload("Transaction.PaymentInfo").Scopes(paginate.Paginate(params.Page, params.Limit)).
		Find(&resp, "user_id = ?", userId).Error
	return
}

func (s *defaultTransaction) FindTransactionDetailByBasketId(ctx context.Context, basketId string) (resp []entity.TransactionDetail, err error) {
	err = s.db.WithContext(ctx).Preload("TransactionBasket").Preload("Product").
		Preload("Product.ProductImg").Find(&resp, "transaction_basket_id = ?", basketId).Error
	return
}

func (s *defaultTransaction) FindTransactionDetailByProductId(ctx context.Context, productId string) (resp entity.TransactionDetail, err error) {
	err = s.db.WithContext(ctx).Take(&resp, "product_id = ?", productId).Error
	return
}

func (s *defaultTransaction) FindTransactionDetailByProductIdAndBasketId(ctx context.Context, productId, basketId string) (resp entity.TransactionDetail, err error) {
	err = s.db.WithContext(ctx).Take(&resp, "product_id = ? AND transaction_basket_id = ?", productId, basketId).Error
	return
}

func (s *defaultTransaction) UpdateTransactionDetail(ctx context.Context, tx *gorm.DB, req *entity.TransactionDetail) (err error) {
	err = tx.WithContext(ctx).Save(req).Error
	return
}

func (s *defaultTransaction) CreateTransactionDetail(ctx context.Context, tx *gorm.DB, req *entity.TransactionDetail) (err error) {
	err = tx.WithContext(ctx).Create(req).Error
	return
}

func (s *defaultTransaction) DeleteTransactionDetail(ctx context.Context, tx *gorm.DB, id int) (err error) {
	err = tx.WithContext(ctx).Delete([]entity.TransactionDetail{}, "id = ?", id).Error
	return
}

func (s *defaultTransaction) CreateTransaction(ctx context.Context, tx *gorm.DB, req *entity.Transaction) (err error) {
	err = tx.WithContext(ctx).Create(req).Error
	return
}

func (s *defaultTransaction) FindTransactionById(ctx context.Context, id string) (resp *entity.Transaction, err error) {
	err = s.db.WithContext(ctx).Preload("PaymentInfo").Preload("Courier").Preload("Courier.CourierInfo").
		Preload("Courier.OriginCity").Preload("Courier.DestinationCity").Preload("TransactionBasket").
		Preload("TransactionBasket.TransactionDetail").Preload("TransactionBasket.TransactionDetail.Product").
		Preload("TransactionBasket.TransactionDetail.Product.ProductImg").Preload("TransactionBasket.User").Take(&resp, "id = ?", id).Error
	return
}

