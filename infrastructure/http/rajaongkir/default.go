package rajaongkir

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/fadilahonespot/cakrawala/utils/logger"
	"github.com/fadilahonespot/library/rest"
)

type wrapper struct {
	client rest.RestClient
}

var address string

func NewWrapper() RajaOngkirWrapper {
	restOptions := rest.Options{
		Address: os.Getenv("RAJA_ONGKIR_HOST"),
		Timeout: time.Duration(10 * time.Second),
		SkipTLS: false,
	}
	client := rest.New(restOptions)
	address = restOptions.Address

	return &wrapper{client: client}
}

func getRequestHeaders(ctx context.Context) (headers http.Header) {

	headers = http.Header{
		"Content-Type": []string{"application/json"},
		"key":          []string{os.Getenv("RAJA_ONGKIR_TOKEN")},
	}

	return
}

func (w *wrapper) GetProvince(ctx context.Context) (resp GetProvinceResponse, err error) {
	path := "/province"

	logger.Info(ctx, "[GetProvince Request]", address+path)

	headers := getRequestHeaders(ctx)
	body, status, err := w.client.Get(ctx, path, headers)
	if err != nil {
		err = fmt.Errorf("[RajaOngkir] GetProvince error: %v", err.Error())
		return
	}

	if status != http.StatusOK {
		err = fmt.Errorf("[RajaOngkir] GetProvince return non 200 http status code. got %d", status)
		return
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		err = fmt.Errorf("[RajaOngkir] Unmarshal Response Error %v", err.Error())
	}

	logger.Info(ctx, "[GetProvince Response]", address+path, resp)

	return
}

func (w *wrapper) GetCity(ctx context.Context, provinceId string) (resp GetCityResponse, err error) {
	path := "/city?province=" + provinceId

	logger.Info(ctx, "[GetProvince Request]", address+path)

	headers := getRequestHeaders(ctx)
	body, status, err := w.client.Get(ctx, path, headers)
	if err != nil {
		err = fmt.Errorf("[RajaOngkir] GetProvince error: %v", err.Error())
		return
	}

	if status != http.StatusOK {
		err = fmt.Errorf("[RajaOngkir] GetProvince return non 200 http status code. got %d", status)
		return
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		err = fmt.Errorf("[RajaOngkir] Unmarshal Response Error %v", err.Error())
	}

	logger.Info(ctx, "[GetProvince Response]", address+path, resp)

	return
}

func (w *wrapper) GetCityAndProvince(ctx context.Context, provinceId, cityId string) (resp GetCityProvinceResponse, err error) {
	path := fmt.Sprintf("/city?id=%v&province=%v", cityId, provinceId)

	logger.Info(ctx, "[GetProvince Request]", address+path)

	headers := getRequestHeaders(ctx)
	body, status, err := w.client.Get(ctx, path, headers)
	if err != nil {
		err = fmt.Errorf("[RajaOngkir] GetProvince error: %v", err.Error())
		return
	}

	if status != http.StatusOK {
		err = fmt.Errorf("[RajaOngkir] GetProvince return non 200 http status code. got %d", status)
		return
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		err = fmt.Errorf("[RajaOngkir] Unmarshal Response Error %v", err.Error())
	}

	logger.Info(ctx, "[GetProvince Response]", address+path, resp)

	return
}

func (w *wrapper) CheckCost(ctx context.Context, req CheckCostRequest) (resp CheckCostResponse, err error) {
	path := "/cost"

	logger.Info(ctx, "[CheckCost Request]", address+path, req)

	headers := getRequestHeaders(ctx)
	body, status, err := w.client.Post(ctx, path, headers, req)
	if err != nil {
		err = fmt.Errorf("[RajaOngkir] CheckCost error: %v", err.Error())
		return
	}

	if status != http.StatusOK {
		err = fmt.Errorf("[RajaOngkir] CheckCost return non 200 http status code. got %d", status)
		return
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		err = fmt.Errorf("[RajaOngkir] Unmarshal Response Error %v", err.Error())
	}

	logger.Info(ctx, "[CheckCost Response]", address+path, resp)

	return
}
