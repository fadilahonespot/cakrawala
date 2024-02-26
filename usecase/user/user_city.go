package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/fadilahonespot/cakrawala/usecase/user/model"
	"github.com/fadilahonespot/cakrawala/utils/errors"
	"github.com/fadilahonespot/cakrawala/utils/logger"
)

func (s *defaultUser) GetCity(ctx context.Context, provinceId string) (resp []model.CityResponse, err error) {
	value := s.cache.GetCity(ctx, provinceId)
	if value != "" {
		err = json.Unmarshal([]byte(value), &resp)
		if err != nil {
			logger.Error(ctx, "failed to unmarshal value: %v", err)
			err = errors.New(http.StatusInternalServerError, "failed get city")
			return
		}
		return
	}

	responData, err := s.rajaOngkirWrapper.GetCity(ctx, provinceId)
	if err != nil {
		logger.Error(ctx, "error get city", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	if responData.RajaOngkir.Status.Code != http.StatusOK {
		logger.Error(ctx, "error get city", responData.RajaOngkir.Status.Description)
		err = errors.New(http.StatusInternalServerError, "failed get city")
		return
	}

	for i := 0; i < len(responData.RajaOngkir.DataCity); i++ {
		data := model.CityResponse{
			CityName: responData.RajaOngkir.DataCity[i].CityName,
			CityId:   responData.RajaOngkir.DataCity[i].CityId,
		}
		resp = append(resp, data)
	}

	respByte, _ := json.Marshal(resp)
	err = s.cache.SetCity(ctx, provinceId, string(respByte))
	if err != nil {
		logger.Error(ctx, "failed set city in redis cache", err.Error())
		err = errors.New(http.StatusInternalServerError, "failed get city")
		return
	}

	return
}
