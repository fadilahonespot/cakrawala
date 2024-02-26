package rajaongkir

type Status struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type DataProvince struct {
	ProvinceId string `json:"province_id"`
	Province   string `json:"province"`
}
type ProvinceTemp struct {
	Status       Status         `json:"status"`
	DataProvince []DataProvince `json:"results"`
}

type GetProvinceResponse struct {
	RajaOngkir ProvinceTemp `json:"rajaongkir"`
}

type DataCity struct {
	CityId     string `json:"city_id"`
	ProvinceId string `json:"province_id"`
	Province   string `json:"province"`
	Type       string `json:"type"`
	CityName   string `json:"city_name"`
	PostalCode string `json:"postal_code"`
}

type CityTemp struct {
	Status   Status     `json:"status"`
	DataCity []DataCity `json:"results"`
}

type GetCityResponse struct {
	RajaOngkir CityTemp `json:"rajaongkir"`
}

type CityProvinceTemp struct {
	Status   Status   `json:"status"`
	DataCity DataCity `json:"results"`
}

type GetCityProvinceResponse struct {
	RajaOngkir CityProvinceTemp `json:"rajaongkir"`
}

type CheckCostRequest struct {
	Origin      int    `json:"origin"`
	Destination int    `json:"destination"`
	Weight      int    `json:"weight"`
	Courier     string `json:"courier"`
}

type DataAddress struct {
	CityID     string `json:"city_id"`
	ProvinceID string `json:"province_id"`
	Province   string `json:"province"`
	Type       string `json:"type"`
	CityName   string `json:"city_name"`
	PostalCode string `json:"postal_code"`
}

type CourierResult struct {
	Code  string     `json:"code"`
	Name  string     `json:"name"`
	Costs []CostList `json:"costs"`
}

type CostList struct {
	Service     string `json:"service"`
	Description string `json:"description"`
	Cost        []Cost `json:"cost"`
}

type Cost struct {
	Value int    `json:"value"`
	Etd   string `json:"etd"`
	Note  string `json:"note"`
}

type CheckResponseList struct {
	Status             Status        `json:"status"`
	OriginDetails      DataAddress   `json:"origin_details"`
	DestinationDetails DataAddress   `json:"destination_details"`
	Results            []CourierResult `json:"results"`
}

type CheckCostResponse struct {
	RajaOngkir CheckResponseList `json:"rajaongkir"`
}
