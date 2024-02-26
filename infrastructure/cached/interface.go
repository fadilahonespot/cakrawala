package cached

import "context"

type RedisClient interface {
	SetEmailVerification(ctx context.Context,  email, code string) (err error)
	GetEmailVerification(ctx context.Context, email string) (code string, err error)
	DeleteEmailVerification(ctx context.Context, email string) (err error) 
	GetCity(ctx context.Context, provinceId string) (data string)
	SetCity(ctx context.Context, provinceId string, value string) (err error)
	GetProvince(ctx context.Context) (data string)
	SetProvince(ctx context.Context, value string) (err error) 
	GetCityProvince(ctx context.Context, provinceId, cityId string) (data string)
	SetCityProvince(ctx context.Context, provinceId, cityId, value string) (err error)
	GetEmailNotif(ctx context.Context, externalId string) (email string)
	SetEmailNotif(ctx context.Context, externalId, value string) (err error)
	GetGenerateText(ctx context.Context, productId string) (resp GenerateText) 
	SetGenerateText(ctx context.Context, productId int, value GenerateText) (err error)
}