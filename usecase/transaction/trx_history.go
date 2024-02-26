package transaction

import (
	"context"
	"net/http"

	"github.com/fadilahonespot/cakrawala/usecase/transaction/model"
	"github.com/fadilahonespot/cakrawala/utils/errors"
	"github.com/fadilahonespot/cakrawala/utils/logger"
	"github.com/fadilahonespot/cakrawala/utils/paginate"
)

func (s *defaultTransaction) GetHistoryTransactio(ctx context.Context, userId string, params paginate.Pagination) (resp []model.HistoryTransactionResponse, count int64, err error) {
	basketData, count, err := s.transactionRepo.FindAllTransactionBasketByUserId(ctx, userId, params)
	if err != nil {
		logger.Error(ctx, "error find data basket transaction", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	for i := 0; i < len(basketData); i++ {
		if len(basketData[i].Transaction) != 0 {
			data := model.HistoryTransactionResponse{
				TransactionId: basketData[i].Transaction[0].ID,
				Xpayment:      basketData[i].Transaction[0].PaymentInfo.XPayment,
				Status:        string(basketData[i].Transaction[0].PaymentInfo.Status),
				CreatedAt:     basketData[i].Transaction[0].PaymentInfo.CreatedAt,
				Amount:        int(basketData[i].Transaction[0].TotalPrice + float64(basketData[i].Transaction[0].Courier.Price)),
			}
			resp = append(resp, data)
		}
	}
	return
}
