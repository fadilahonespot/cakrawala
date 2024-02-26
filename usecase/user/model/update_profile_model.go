package model

type UpdateProfileRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
	CityID      int    `json:"cityId"`
	ProvinceID  int    `json:"provinceId"`
	PostalCode  string `json:"postalCode"`
}
