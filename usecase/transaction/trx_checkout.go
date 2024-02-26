package transaction

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/fadilahonespot/cakrawala/domain/entity"
	"github.com/fadilahonespot/cakrawala/infrastructure/http/xendit"
	"github.com/fadilahonespot/cakrawala/usecase/transaction/model"
	"github.com/fadilahonespot/cakrawala/utils"
	"github.com/fadilahonespot/cakrawala/utils/constrans"
	"github.com/fadilahonespot/cakrawala/utils/errors"
	"github.com/fadilahonespot/cakrawala/utils/logger"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

func (s *defaultTransaction) Checkout(ctx context.Context, userId string, req model.CheckoutRequest) (resp model.CheckoutResponse, err error) {
	courierInfoData, err := s.courierInfoRepo.FindByCode(ctx, req.CourierCode)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Error(ctx, "courier not found")
			err = errors.New(http.StatusNotFound, "Courier not found")
			return
		}

		logger.Error(ctx, "faled find data", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	userData, err := s.userRepo.FindByID(ctx, cast.ToInt(userId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Error(ctx, "user not found")
			err = errors.New(http.StatusNotFound, "user not found")
			return
		}
		logger.Error(ctx, "failed to find user", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	if userData.EmailStatus == entity.Unverify {
		err = errors.New(http.StatusForbidden, "email status is not verified")
		return
	}
	basketData, err := s.transactionRepo.FindTransactionBasketByUserIdStatusPending(ctx, userId)
	if err != nil {
		logger.Error(ctx, "basket transaction not found", err.Error())
		err = errors.New(http.StatusNotFound, "There are no items in the basket yet")
		return
	}
	trxDetail, err := s.transactionRepo.FindTransactionDetailByBasketId(ctx, strconv.Itoa(basketData.ID))
	if err != nil {
		logger.Error(ctx, "transaction detail not found", err.Error())
		err = errors.New(http.StatusNotFound, "There are no items in the basket yet")
		return
	}

	var totalAmount float64
	var totalWeight int
	var totalQuality int
	for i := 0; i < len(trxDetail); i++ {
		totalAmount += trxDetail[i].Price
		totalWeight += trxDetail[i].Weight
		totalQuality += trxDetail[i].Qty
	}
	dataShipping, err := s.getCostShipping(ctx, constrans.ShopCity, userData.CityID, totalWeight, req.CourierCode)
	if err != nil {
		return
	}

	var courierService model.CourierService
	for k := 0; k < len(dataShipping.CourierService); k++ {
		if strings.EqualFold(dataShipping.CourierService[k].ServiceName, req.CourierService) {
			data := model.CourierService{
				ServiceName: dataShipping.CourierService[k].ServiceName,
				Description: dataShipping.CourierService[k].Description,
				Price:       dataShipping.CourierService[k].Price,
				Etd:         dataShipping.CourierService[k].Etd,
				Note:        dataShipping.CourierService[k].Note,
			}
			courierService = data
			break
		}
	}

	if courierService.ServiceName == "" {
		logger.Error(ctx, "courier service not found")
		err = errors.New(http.StatusNotFound, "courier service not found")
		return
	}

	tx := s.transactionRepo.BeginTrans(ctx)
	cityData, _ := s.addressRepo.GetCity(ctx, fmt.Sprintf("%v", constrans.ShopCity))
	if cityData.ID == 0 {
		data := entity.City{
			ID:   constrans.ShopCity,
			Name: dataShipping.DestinationCity,
		}
		err = s.addressRepo.UpdateCity(ctx, &data)
		if err != nil {
			tx.Rollback()
			err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			logger.Error(ctx, "failed to update city", err.Error())
			return
		}
	}
	courierReq := entity.Courier{
		OriginCityID:      constrans.ShopCity,
		DestinationCityID: userData.CityID,
		CourierInfoID:     courierInfoData.ID,
		Service:           courierService.ServiceName,
		Description:       courierService.Description,
		Price:             courierService.Price,
		ETD:               courierService.Etd,
		Note:              courierService.Note,
	}
	err = s.courierRepo.Create(ctx, tx, &courierReq)
	if err != nil {
		tx.Rollback()
		logger.Error(ctx, "error creating courier", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	reqVa := xendit.CreateVirtualAccountRequest{
		ExternalID:     utils.GenerateExternalId(15),
		BankCode:       req.BankCode,
		Name:           userData.Name,
		IsSingleUse:    true,
		IsClosed:       true,
		ExpectedAmount: int(totalAmount) + courierService.Price,
		ExpirationDate: time.Now().Add(time.Duration(24) * time.Hour),
	}
	respVa, err := s.xenditWrapper.CreateVirtualAccount(ctx, reqVa)
	if err != nil {
		tx.Rollback()
		logger.Error(ctx, "failed creating virtual account", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	paymentReq := entity.PaymentInfo{
		XPayment:       respVa.ExternalID,
		Status:         entity.Unpaid,
		BankCode:       req.BankCode,
		AccountNumber:  respVa.AccountNumber,
		Amount:         totalAmount,
		ExpirationDate: respVa.ExpirationDate,
	}
	err = s.paymentInfoRepo.Create(ctx, &paymentReq)
	if err != nil {
		tx.Rollback()
		logger.Error(ctx, "error creating payment info", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	trxReq := entity.Transaction{
		TransactionBasketID: basketData.ID,
		TotalQty:            totalQuality,
		TotalPrice:          totalAmount,
		TotalWeight:         totalWeight,
		PaymentInfoID:       paymentReq.ID,
		CourierID:           courierReq.ID,
	}
	err = s.transactionRepo.CreateTransaction(ctx, tx, &trxReq)
	if err != nil {
		tx.Rollback()
		logger.Error(ctx, "error creating transaction", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	reqBasket := entity.TransactionBasket{
		ID:           basketData.ID,
		UserID:       basketData.UserID,
		BasketStatus: entity.BasketCompleted,
		CreatedAt:    basketData.CreatedAt,
	}
	err = s.transactionRepo.UpdateTransactionBasket(ctx, tx, &reqBasket)
	if err != nil {
		tx.Rollback()
		logger.Error(ctx, "error update basket data", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	resp = model.CheckoutResponse{
		TransactionID:      trxReq.ID,
		XPayment:           respVa.ExternalID,
		BankCode:           respVa.BankCode,
		VirtualAccount:     respVa.AccountNumber,
		VirtualAccountName: respVa.Name,
		ShippingCost:       courierService.Price,
		ProductPrice:       int(totalAmount),
		TotalPayment:       respVa.ExpectedAmount,
		ExpiredDate:        respVa.ExpirationDate,
	}

	if userData.EmailStatus == entity.Verify {
		go s.sendPaymentNotif(ctx, userData.Email, resp)
		s.cache.SetEmailNotif(ctx, respVa.ExternalID, userData.Email)
	}

	tx.Commit()
	return
}

func (s *defaultTransaction) sendPaymentNotif(ctx context.Context, email string, req model.CheckoutResponse) {
	body, err := generatePaymentNotifBody(req)
	if err != nil {
		logger.Error(ctx, "error generate notif", err.Error())
	}
	err = s.mailjetWrapper.SendEmail(ctx, email, "Notifikasi Pembayaran", body)
	if err != nil {
		logger.Error(ctx, "error send email notif payment", err.Error())
	}
}
