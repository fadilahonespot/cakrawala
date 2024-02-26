package xendit

import "context"

type XenditWrapper interface {
	GetBank(ctx context.Context) (resp []GetBankResponse, err error)
	CreateVirtualAccount(ctx context.Context, req CreateVirtualAccountRequest) (resp CreateViartualAccountResponse, err error)
	CheckVirtualAccount(ctx context.Context, id string) (resp CreateViartualAccountResponse, err error)
	CheckPayment(ctx context.Context, paymentId string) (resp CheckPaymentResponse, err error)
}