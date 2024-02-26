package user

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/fadilahonespot/cakrawala/domain/entity"
	"github.com/fadilahonespot/cakrawala/infrastructure/http/rajaongkir"
	"github.com/fadilahonespot/cakrawala/usecase/user/model"
	"github.com/fadilahonespot/cakrawala/utils"
	"github.com/fadilahonespot/cakrawala/utils/bcrypt"
	"github.com/fadilahonespot/cakrawala/utils/constrans"
	"github.com/fadilahonespot/cakrawala/utils/errors"
	"github.com/fadilahonespot/cakrawala/utils/logger"
	"github.com/spf13/cast"
)

func (s *defaultUser) Register(ctx context.Context, req model.RegisterRequest) (err error) {
	phoneNumber := utils.CorrectPhoneNumber(req.PhoneNumber)
	req.PhoneNumber = phoneNumber
	_, err = s.userRepo.FindByEmail(ctx, req.Email)
	if err == nil {
		logger.Error(ctx, "email already registered")
		err = errors.New(http.StatusConflict, "email already exists")
		return
	}

	_, err = s.GetCityProvince(ctx, req.ProvinceID, req.CityID)
	if err != nil {
		return
	}

	passBcrypt := bcrypt.HashSHA256(constrans.SecretPassword, req.Password)
	userData := entity.User{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		EmailStatus: entity.Unverify,
		Password:    passBcrypt,
		Address:     req.Address,
		CityID:      req.CityID,
		ProvinceID:  req.ProvinceID,
		Role:        entity.Customer, // default is customer
		PostalCode:  req.PostalCode,
	}

	err = s.userRepo.Create(ctx, &userData)
	if err != nil {
		logger.Error(ctx, "failed to create user", err.Error())
		err = errors.New(http.StatusInternalServerError, "failed to create user")
		return
	}

	go s.sendEmailVerification(ctx, userData.Name, req.Email)
	return
}

func (s *defaultUser) GetCityProvince(ctx context.Context, provinceId, cityId int) (resp rajaongkir.GetCityProvinceResponse, err error) {
	var dataCity rajaongkir.GetCityProvinceResponse
	value := s.cache.GetCityProvince(ctx, strconv.Itoa(provinceId), strconv.Itoa(cityId))
	switch true {
	case value == "":
		dataCity, err = s.rajaOngkirWrapper.GetCityAndProvince(ctx, strconv.Itoa(provinceId), strconv.Itoa(cityId))
		if err != nil {
			logger.Error(ctx, "error get city and province", err.Error())
			err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}

		if dataCity.RajaOngkir.DataCity.CityId == "" {
			logger.Error(ctx, "city or province not found")
			err = errors.New(http.StatusNotFound, "city or province not found")
			return
		}

		dataCitybyte, _ := json.Marshal(dataCity)
		s.cache.SetCityProvince(ctx, strconv.Itoa(provinceId), strconv.Itoa(cityId), string(dataCitybyte))
		if err != nil {
			logger.Error(ctx, "failed set province in redis cache", err.Error())
			err = errors.New(http.StatusInternalServerError, "failed get province")
			return
		}

	default:
		err = json.Unmarshal([]byte(value), &dataCity)
		if err != nil {
			logger.Error(ctx, "failed unmarshal value", err.Error())
			err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
	}

	reqCity := entity.City{
		ID:        cast.ToInt(dataCity.RajaOngkir.DataCity.CityId),
		Name:      dataCity.RajaOngkir.DataCity.CityName,
		CreatedAt: time.Now(),
	}
	err = s.addressRepo.UpdateCity(ctx, &reqCity)
	if err != nil {
		logger.Error(ctx, "failed update city", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	reqProvince := entity.Province{
		ID:        cast.ToInt(dataCity.RajaOngkir.DataCity.ProvinceId),
		Name:      dataCity.RajaOngkir.DataCity.Province,
		CreatedAt: time.Now(),
	}
	err = s.addressRepo.UpdateProvince(ctx, &reqProvince)
	if err != nil {
		logger.Error(ctx, "failed update province", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	resp = dataCity
	return
}
