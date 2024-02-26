package user

import (
	"context"
	"net/http"

	"github.com/fadilahonespot/cakrawala/usecase/user/model"
	"github.com/fadilahonespot/cakrawala/utils"
	"github.com/fadilahonespot/cakrawala/utils/errors"
	"github.com/fadilahonespot/cakrawala/utils/logger"
	"gorm.io/gorm"
)

func (s *defaultUser) UpdateProfile(ctx context.Context, userId int, req model.UpdateProfileRequest) (err error) {
	userData, err := s.userRepo.FindByID(ctx, userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Error(ctx, "user not found", err.Error())
			err = errors.New(http.StatusNotFound, "user not found")
			return
		}
		logger.Error(ctx, "error find data user", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	_, err = s.GetCityProvince(ctx, req.ProvinceID, req.CityID)
	if err != nil {
		return
	}

	userData.PhoneNumber = utils.CorrectPhoneNumber(req.PhoneNumber)
	userData.Address = req.Address
	userData.Name = req.Name
	userData.CityID = req.CityID
	userData.PostalCode = req.PostalCode
	userData.ProvinceID = req.ProvinceID

	err = s.userRepo.Update(ctx, userData)
	if err != nil {
		logger.Error(ctx, "error update data user", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	return
}
