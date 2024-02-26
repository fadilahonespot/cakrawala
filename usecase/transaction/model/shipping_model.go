package model

type CheckShippingRequest struct {
	CourierCode string `json:"courierCode"`
	Weight      int    `json:"weight"`
}

type CheckShippingResponse struct {
	CourierCode         string           `json:"courierCode"`
	CourierName         string           `json:"courierName"`
	OriginProvince      string           `json:"originProvince"`
	OriginCity          string           `json:"originCity"`
	DestinationProvince string           `json:"destinationProvince"`
	DestinationCity     string           `json:"destinationCity"`
	CourierService      []CourierService `json:"courierService"`
}

type CheckShippingData struct {
	CheckShippingResponse
	OriginCityId          string
	OriginProvinceId      string
	DestinationCityId     string
	DestinationProvinceId string
}

type CourierService struct {
	ServiceName string `json:"serviceName"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Etd         string `json:"etd"`
	Note        string `json:"note"`
}
