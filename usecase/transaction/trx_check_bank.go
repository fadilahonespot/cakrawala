package transaction

import (
	"context"
	"net/http"

	"github.com/fadilahonespot/cakrawala/usecase/transaction/model"
	"github.com/fadilahonespot/cakrawala/utils/errors"
	"github.com/fadilahonespot/cakrawala/utils/logger"
)

func (s *defaultTransaction) CheckAvailableBank(ctx context.Context) (resp []model.CheckBankResponse, err error) {
	dataBank, err := s.xenditWrapper.GetBank(ctx)
	if err != nil {
		logger.Error(ctx, "error get bank", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	for i := 0; i < len(dataBank); i++ {
		if dataBank[i].IsActivated  && dataBank[i].Country == "ID" && dataBank[i].Currency == "IDR" {
			resp = append(resp, model.CheckBankResponse{
				Name: dataBank[i].Name,
				Code: dataBank[i].Code,
			})
		}
	}

	return
}
