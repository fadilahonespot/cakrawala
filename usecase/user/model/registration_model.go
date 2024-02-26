package model

type RegisterRequest struct {
	Email       string `json:"email" validate:"email,required"`
	Password    string `json:"password" validate:"required,min=6"`
	Name        string `json:"name" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required,min=10,max=15"`
	Address     string `json:"address" validate:"required"`
	RoleID      int    `json:"roleId"`
	CityID      int    `json:"cityId" validate:"required"`
	ProvinceID  int    `json:"provinceId" validate:"required"`
	PostalCode  string `json:"postalCode" validate:"required"`
}
