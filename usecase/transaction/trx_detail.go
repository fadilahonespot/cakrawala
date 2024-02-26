package transaction

import (
	"context"
	"net/http"

	"github.com/fadilahonespot/cakrawala/usecase/transaction/model"
	"github.com/fadilahonespot/cakrawala/utils/errors"
	"github.com/fadilahonespot/cakrawala/utils/logger"
	"gorm.io/gorm"
)

func (s *defaultTransaction) GetDetailTransaction(ctx context.Context, transactionId string) (resp model.DetailTransactionResponse, err error) {
	trxData, err := s.transactionRepo.FindTransactionById(ctx, transactionId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Error(ctx, "transaction not found", err.Error())
			err = errors.New(http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}
		logger.Error(ctx, "error find transaction", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	adminFree := 0
	resp = model.DetailTransactionResponse{
		UserName:        trxData.TransactionBasket.User.Name,
		UserAddress:     trxData.TransactionBasket.User.Address,
		UserPhone:       trxData.TransactionBasket.User.PhoneNumber,
		UserEmail:       trxData.TransactionBasket.User.Email,
		TransactionId:   trxData.ID,
		Xpayment:        trxData.PaymentInfo.XPayment,
		Status:          string(trxData.PaymentInfo.Status),
		Bank:            trxData.PaymentInfo.BankCode,
		VirtualAccount:  trxData.PaymentInfo.AccountNumber,
		AdminFree:       adminFree,
		ShippingCost:    trxData.Courier.Price,
		ProductAmount:   int(trxData.TotalPrice),
		TotalPayment:    trxData.Courier.Price + int(trxData.TotalPrice) + adminFree,
		ProductQuantity: trxData.TotalQty,
		ProductWeight:   trxData.TotalWeight,
		CourierName:     trxData.Courier.CourierInfo.Code,
		CourierService:  trxData.Courier.Service,
		OriginCity:      trxData.Courier.OriginCity.Name,
		DestinationCity: trxData.Courier.DestinationCity.Name,
		Etd:             trxData.Courier.ETD,
		CreatedAt:       trxData.CreatedAt,
		ExpiredAt:       trxData.PaymentInfo.ExpirationDate,
	}

	items := trxData.TransactionBasket.TransactionDetail
	for i := 0; i < len(items); i++ {
		
		item := model.TrxItem{
			ProductID:   items[i].Product.ID,
			ProductName: items[i].Product.Name,
			Price:       items[i].Product.Price,
			Weight:      items[i].Product.Weight,
			Quantity:    items[i].Qty,
			TotalPrice:  items[i].Price,
			TotalWeight: float64(items[i].Weight),
		}
		if len(items[i].Product.ProductImg) != 0 {
			item.ProductImg = items[i].Product.ProductImg[0].Image
		}
		resp.Items = append(resp.Items, item)
	}

	return
}
