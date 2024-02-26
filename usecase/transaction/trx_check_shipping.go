package transaction

import (
	"context"
	"net/http"

	"github.com/fadilahonespot/cakrawala/domain/entity"
	"github.com/fadilahonespot/cakrawala/infrastructure/http/rajaongkir"
	"github.com/fadilahonespot/cakrawala/usecase/transaction/model"
	"github.com/fadilahonespot/cakrawala/utils/constrans"
	"github.com/fadilahonespot/cakrawala/utils/errors"
	"github.com/fadilahonespot/cakrawala/utils/logger"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

func (s *defaultTransaction) CheckShipping(ctx context.Context, userId string, req model.CheckShippingRequest) (resp model.CheckShippingResponse, err error) {
	_, err = s.courierInfoRepo.FindByCode(ctx, req.CourierCode)
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

	data, err := s.getCostShipping(ctx, constrans.ShopCity, userData.CityID, req.Weight, req.CourierCode)
	if err != nil {
		return
	}

	originCity := entity.City{
		ID: cast.ToInt(data.OriginCityId),
		Name: data.OriginCity,
	}
	err = s.addressRepo.UpdateCity(ctx, &originCity)
	if err != nil {
		logger.Error(ctx, "failed to update city", err.Error())
	}

	originProvince := entity.Province{
		ID: cast.ToInt(data.OriginProvinceId),
		Name: data.OriginProvince,
	}
	err = s.addressRepo.UpdateProvince(ctx, &originProvince)
	if err != nil {
		logger.Error(ctx, "failed to update province", err.Error())
	}

	resp = data.CheckShippingResponse
	return
}

func (s *defaultTransaction) getCostShipping(ctx context.Context, origin, destination, weight int, courierCode string) (resp model.CheckShippingData, err error) {
	reqCost := rajaongkir.CheckCostRequest{
		Origin:      origin,
		Destination: destination,
		Weight:      weight,
		Courier:     courierCode,
	}
	data, err := s.rajaongkirWrapper.CheckCost(ctx, reqCost)
	if err != nil {
		logger.Error(ctx, "failed to check cost", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	resp = model.CheckShippingData{
		CheckShippingResponse: model.CheckShippingResponse{
			CourierCode:         courierCode,
			OriginCity:          data.RajaOngkir.OriginDetails.CityName,
			OriginProvince:      data.RajaOngkir.OriginDetails.Province,
			DestinationCity:     data.RajaOngkir.DestinationDetails.CityName,
			DestinationProvince: data.RajaOngkir.DestinationDetails.Province,
		},
		OriginCityId:          data.RajaOngkir.OriginDetails.CityID,
		OriginProvinceId:      data.RajaOngkir.OriginDetails.ProvinceID,
		DestinationCityId:     data.RajaOngkir.DestinationDetails.CityID,
		DestinationProvinceId: data.RajaOngkir.DestinationDetails.ProvinceID,
	}

	if len(data.RajaOngkir.Results) != 0 {
		resp.CourierName = data.RajaOngkir.Results[0].Name

		costs := data.RajaOngkir.Results[0].Costs
		for i := 0; i < len(costs); i++ {
			temp := model.CourierService{
				ServiceName: costs[i].Service,
				Description: costs[i].Description,
			}

			if len(costs[i].Cost) != 0 {
				temp.Price = costs[i].Cost[0].Value
				temp.Etd = costs[i].Cost[0].Etd
				temp.Note = costs[i].Cost[0].Note
			}
			resp.CourierService = append(resp.CourierService, temp)
		}
	}
	return
}
