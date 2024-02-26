package user

import (
	"context"
	"net/http"
	"time"

	"github.com/fadilahonespot/cakrawala/usecase/user/model"
	"github.com/fadilahonespot/cakrawala/utils/bcrypt"
	"github.com/fadilahonespot/cakrawala/utils/constrans"
	"github.com/fadilahonespot/cakrawala/utils/errors"
	"github.com/fadilahonespot/cakrawala/utils/logger"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

func (u *defaultUser) Login(ctx context.Context, req model.LoginRequest) (resp model.LoginResponse, err error) {
	userData, err := u.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Error(ctx, "user not found")
			err = errors.New(http.StatusNotFound, "User not found")
			return
		}

		logger.Error(ctx, "faled find data", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	passBcrypt := bcrypt.HashSHA256(constrans.SecretPassword, req.Password)
	if passBcrypt != userData.Password {
		logger.Error(ctx, "user wrong password")
		err = errors.New(http.StatusUnauthorized, "wrong password")
		return
	}

	claims := jwt.MapClaims{
		"userId": userData.ID,
		"role":   userData.Role,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	tokenClams := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClams.SignedString([]byte(constrans.JwtSecret))
	if err != nil {
		logger.Error(ctx, "failed generate token")
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	resp = model.LoginResponse{
		Name:   userData.Name,
		Role:   string(userData.Role),
		UserId: userData.ID,
		Token:  token,
	}
	return
}
