package transaction

import (
	"context"
	"net/http"

	"github.com/fadilahonespot/cakrawala/usecase/transaction/model"
	"github.com/fadilahonespot/cakrawala/utils/errors"
	"github.com/fadilahonespot/cakrawala/utils/logger"
)

func (s *defaultTransaction) GetCourierInfo(ctx context.Context) (resp []model.CourierInfoResponse, err error) {
	data, err := s.courierInfoRepo.FindAll(ctx)
	if err != nil {
		logger.Error(ctx, "failed find courier info")
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	for i := 0; i < len(data); i++ {
		resp = append(resp, model.CourierInfoResponse{
			Name:  data[i].Name,
			Code:  data[i].Code,
			Image: data[i].Image,
		})
	}
	return
}
