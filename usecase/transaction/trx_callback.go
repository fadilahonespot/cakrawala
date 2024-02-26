package transaction

import (
	"context"
	"net/http"

	"github.com/fadilahonespot/cakrawala/domain/entity"
	"github.com/fadilahonespot/cakrawala/usecase/transaction/model"
	"github.com/fadilahonespot/cakrawala/utils/errors"
	"github.com/fadilahonespot/cakrawala/utils/logger"
)

func (s *defaultTransaction) CallbackTransaction(ctx context.Context, webhookId string, req model.CallbackRequest) (err error) {
	paymentData, err := s.paymentInfoRepo.FindByXpayment(ctx, req.ExternalID)
	if err != nil {
		logger.Error(ctx, "failed to find payment info", err.Error())
		err = errors.New(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	xenditResp, err := s.xenditWrapper.CheckPayment(ctx, req.PaymentID)
	if err != nil {
		logger.Error(ctx, "failed to check payment", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	if xenditResp.CallbackVirtualAccountID != "" {
		paymentData.Status = entity.Paid
		paymentData.PaymentTime = &xenditResp.TransactionTimestamp
	} else {
		paymentData.Status = entity.Failed
	}
	
	err = s.paymentInfoRepo.Update(ctx, paymentData)
	if err != nil {
		logger.Error(ctx, "failed to update payment info", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	if paymentData.Status == entity.Paid {
		email := s.cache.GetEmailNotif(ctx, xenditResp.ExternalID)
		if email != "" {
			body, err := generatePaymentSuccessBody(xenditResp, *paymentData)
			if err != nil {
				logger.Error(ctx, "error generate notif", err.Error())
			}
			err = s.mailjetWrapper.SendEmail(ctx, email, "Notifikasi Pembayaran", body)
			if err != nil {
				logger.Error(ctx, "error send email notif payment", err.Error())
			}
		}
	}

	return
}
