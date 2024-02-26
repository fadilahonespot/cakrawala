package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/fadilahonespot/cakrawala/usecase/user/model"
	"github.com/fadilahonespot/cakrawala/utils/errors"
	"github.com/fadilahonespot/cakrawala/utils/logger"
)

func (s *defaultUser) GetProvince(ctx context.Context) (resp []model.ProvinceResponse, err error) {
	value := s.cache.GetProvince(ctx)
	if value != "" {
		err = json.Unmarshal([]byte(value), &resp)
		if err != nil {
			logger.Error(ctx, "failed to unmarshal value: %v", err)
			err = errors.New(http.StatusInternalServerError, "failed get province")
			return
		}
		return
	}

	responData, err := s.rajaOngkirWrapper.GetProvince(ctx)
	if err != nil {
		logger.Error(ctx, "error get province", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	if responData.RajaOngkir.Status.Code != http.StatusOK {
		logger.Error(ctx, "error get province", responData.RajaOngkir.Status.Description)
		err = errors.New(http.StatusInternalServerError, "failed get province")
		return
	}

	for i := 0; i < len(responData.RajaOngkir.DataProvince); i++ {
		data := model.ProvinceResponse{
			ProvinceName: responData.RajaOngkir.DataProvince[i].Province,
			ProvinceId:   responData.RajaOngkir.DataProvince[i].ProvinceId,
		}
		resp = append(resp, data)
	}

	respByte, _ := json.Marshal(resp)
	err = s.cache.SetProvince(ctx, string(respByte))
	if err != nil {
		logger.Error(ctx, "failed set province in redis cache", err.Error())
		err = errors.New(http.StatusInternalServerError, "failed get province")
		return
	}

	return
}
