package transaction

import (
	"context"

	"github.com/fadilahonespot/cakrawala/usecase/transaction/model"
	"github.com/fadilahonespot/cakrawala/utils/paginate"
)

type TransactionService interface {
	CheckAvailableBank(ctx context.Context) (resp []model.CheckBankResponse, err error)
	GetCourierInfo(ctx context.Context) (resp []model.CourierInfoResponse, err error)
	CheckShipping(ctx context.Context, userId string, req model.CheckShippingRequest) (resp model.CheckShippingResponse, err error)
	GetAllProductBasket(ctx context.Context, userId string) (resp []model.BasketResponse, err error)
	AddProductBasket(ctx context.Context, userId string, req model.AddBasketRequest) (err error)
	DeleteProductBasket(ctx context.Context, userId, productId string) (err error)
	Checkout(ctx context.Context, userId string, req model.CheckoutRequest) (resp model.CheckoutResponse, err error)
	CallbackTransaction(ctx context.Context, webhookId string, req model.CallbackRequest) (err error)
	GetHistoryTransactio(ctx context.Context, userId string,params paginate.Pagination) (resp []model.HistoryTransactionResponse, count int64, err error)
	GetDetailTransaction(ctx context.Context, transactionId string) (resp model.DetailTransactionResponse, err error)
}