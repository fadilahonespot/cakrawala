package model

type GetProfileRequest struct {
	UserId   string `json:"userId"`
	RoleId   string `json:"roleId"`
	RoleName string `json:"roleName"`
}

type GetProfileResponse struct {
	Email        string `json:"email"`
	EmailStatus  string `json:"emailStatus"`
	Name         string `json:"name"`
	PhoneNumber  string `json:"phoneNumber"`
	Address      string `json:"address"`
	RoleName     string `json:"roleName"`
	CityID       int    `json:"cityId"`
	ProvinceID   int    `json:"provinceId"`
	CityName     string `json:"cityName"`
	ProvinceName string `json:"provinceName"`
	PostalCode   string `json:"postalCode"`
}
