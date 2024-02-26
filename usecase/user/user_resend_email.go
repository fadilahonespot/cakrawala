package user

import (
	"context"
	"net/http"

	"github.com/fadilahonespot/cakrawala/domain/entity"
	"github.com/fadilahonespot/cakrawala/utils/errors"
	"github.com/fadilahonespot/cakrawala/utils/logger"
	"github.com/spf13/cast"
)

func (s *defaultUser) ResendEmail(ctx context.Context, userId string) (err error) {
	userData, err := s.userRepo.FindByID(ctx, cast.ToInt(userId))
	if err != nil {
		logger.Error(ctx, "user not found")
		err = errors.New(http.StatusNotFound, "user not found exists")
		return
	}

	if userData.EmailStatus == entity.Verify {
		logger.Error(ctx, "email already verified")
        err = errors.New(http.StatusBadRequest, "email already verified")
        return
	}

	err = s.sendEmailVerification(ctx, userData.Name, userData.Email)
	if err != nil {
		logger.Error(ctx, "failed send email", err.Error())
        err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
        return
	}

	return
}