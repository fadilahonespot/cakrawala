package user

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/fadilahonespot/cakrawala/domain/repository"
	"github.com/fadilahonespot/cakrawala/infrastructure/cached"
	"github.com/fadilahonespot/cakrawala/infrastructure/http/mailjet"
	"github.com/fadilahonespot/cakrawala/infrastructure/http/rajaongkir"
	"github.com/fadilahonespot/cakrawala/utils"
	"github.com/fadilahonespot/cakrawala/utils/errors"
	"github.com/fadilahonespot/cakrawala/utils/logger"
)

type defaultUser struct {
	userRepo          repository.UserRepository
	rajaOngkirWrapper rajaongkir.RajaOngkirWrapper
	cache             cached.RedisClient
	addressRepo       repository.AddressRepository
	mailjetWrapper    mailjet.MailjetWrapper
}

func SetupUserService() *defaultUser {
	return &defaultUser{}
}

func (u *defaultUser) SetUserRepo(userRepo repository.UserRepository) *defaultUser {
	u.userRepo = userRepo
	return u
}

func (u *defaultUser) SetRajaOngkirWrapper(wrapper rajaongkir.RajaOngkirWrapper) *defaultUser {
	u.rajaOngkirWrapper = wrapper
	return u
}

func (u *defaultUser) SetRedisClient(client cached.RedisClient) *defaultUser {
	u.cache = client
	return u
}

func (u *defaultUser) SetAddressRepo(addressRepo repository.AddressRepository) *defaultUser {
	u.addressRepo = addressRepo
	return u
}

func (u *defaultUser) SetMailjetWrapper(wrapper mailjet.MailjetWrapper) *defaultUser {
	u.mailjetWrapper = wrapper
	return u
}

func (u *defaultUser) Validate() UserService {
	if u.userRepo == nil {
		panic("user repo is nil")
	}

	if u.rajaOngkirWrapper == nil {
		panic("raja ongkir warpper is nil")
	}

	if u.cache == nil {
		panic("redis client is nil")
	}

	if u.addressRepo == nil {
		panic("address repo is nil")
	}

	if u.mailjetWrapper == nil {
		panic("mailjet wrapper is nil")
	}

	return u
}

func (s *defaultUser) sendEmailVerification(ctx context.Context, name, email string) (err error) {
	verificatiCode := utils.GenerateRandomString(100)
	err = s.cache.SetEmailVerification(ctx, email, verificatiCode)
	if err != nil {
		logger.Error(ctx, "failed set verification email code in redis cache", err.Error())
		err = errors.New(http.StatusInternalServerError, "failed get province")
		return
	}

	url := fmt.Sprintf("%v/external/user/verification?email=%v&id=%v", os.Getenv("VERIFICATION_EMAIL_HOST"), email, verificatiCode)
	emailBody, err := utils.GenerateEmailVerificationBody(name, url)
	if err != nil {
		logger.Error(ctx, "failed generate email verification body", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	err = s.mailjetWrapper.SendEmail(ctx, email, "Konfirmasi Alamat Email Kamu", emailBody)
	if err != nil {
		logger.Error(ctx, "failed send email", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	return
}
