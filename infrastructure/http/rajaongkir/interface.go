package rajaongkir

import "context"

type RajaOngkirWrapper interface {
	GetProvince(ctx context.Context) (resp GetProvinceResponse, err error)
	GetCity(ctx context.Context, provinceId string) (resp GetCityResponse, err error)
	GetCityAndProvince(ctx context.Context, provinceId, cityId string) (resp GetCityProvinceResponse, err error)
	CheckCost(ctx context.Context, req CheckCostRequest) (resp CheckCostResponse, err error)
}