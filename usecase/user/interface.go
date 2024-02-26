package user

import (
	"context"

	"github.com/fadilahonespot/cakrawala/usecase/user/model"
)

type UserService interface {
	Login(ctx context.Context, req model.LoginRequest) (resp model.LoginResponse, err error) 
	Register(ctx context.Context, req model.RegisterRequest) (err error)
	GetProvince(ctx context.Context) (resp []model.ProvinceResponse, err error)
	GetCity(ctx context.Context, provinceId string) (resp []model.CityResponse, err error)
	VerificationEmail(ctx context.Context, req model.VerificationRequest) (err error)
	ResendEmail(ctx context.Context, userId string) (err error)
	GetProfile(ctx context.Context, req model.GetProfileRequest) (resp model.GetProfileResponse, err error)
	UpdateProfile(ctx context.Context, userId int, req model.UpdateProfileRequest) (err error)
}