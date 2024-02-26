package user

import (
	"context"
	"net/http"

	"github.com/fadilahonespot/cakrawala/usecase/user/model"
	"github.com/fadilahonespot/cakrawala/utils/errors"
	"github.com/fadilahonespot/cakrawala/utils/logger"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

func (s *defaultUser) GetProfile(ctx context.Context, req model.GetProfileRequest) (resp model.GetProfileResponse, err error) {
	userData, err := s.userRepo.FindByID(ctx, cast.ToInt(req.UserId))
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

	cityData, err := s.GetCityProvince(ctx, userData.ProvinceID, userData.CityID)
	if err != nil {
		return
	}

	resp = model.GetProfileResponse{
		Email:        userData.Email,
		EmailStatus:  string(userData.EmailStatus),
		Name:         userData.Name,
		PhoneNumber:  userData.PhoneNumber,
		Address:      userData.Address,
		RoleName:     req.RoleName,
		CityID:       cast.ToInt(cityData.RajaOngkir.DataCity.CityId),
		ProvinceID:  cast.ToInt( cityData.RajaOngkir.DataCity.ProvinceId),
		CityName:     cityData.RajaOngkir.DataCity.CityName,
		ProvinceName: cityData.RajaOngkir.DataCity.Province,
		PostalCode:   userData.PostalCode,
	}
	return
}
