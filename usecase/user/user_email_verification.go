package user

import (
	"context"
	"net/http"

	"github.com/fadilahonespot/cakrawala/domain/entity"
	"github.com/fadilahonespot/cakrawala/usecase/user/model"
	"github.com/fadilahonespot/cakrawala/utils"
	"github.com/fadilahonespot/cakrawala/utils/errors"
	"github.com/fadilahonespot/cakrawala/utils/logger"
)

func (s *defaultUser) VerificationEmail(ctx context.Context, req model.VerificationRequest) (err error) {
	codeValue, _ := s.cache.GetEmailVerification(ctx, req.Email)
	if codeValue == "" {
		logger.Error(ctx, "email not found in redis cache")
        err = errors.New(http.StatusNotFound, "link verification not valid")
        return
	}

	if codeValue != req.Id {
		logger.Error(ctx, "id verification not same as email")
        err = errors.New(http.StatusNotFound, "link verification not valid")
        return
	}

	userData, err := s.userRepo.FindByEmail(ctx, req.Email)
    if err != nil {
        logger.Error(ctx, "email verification not found", err.Error())
        err = errors.New(http.StatusNotFound, "link verification not valid")
        return
    }

	if userData.EmailStatus == entity.Verify {
		logger.Error(ctx, "email already verified")
        err = errors.New(http.StatusConflict, "email already verified")
        return
	}

	userData.EmailStatus = entity.Verify
	err = s.userRepo.Update(ctx, userData)
	if err != nil {
		logger.Error(ctx, "failed to update user", err.Error())
        err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
        return
	}

	emailBody := utils.GenerateEmailVerificationSuccesBody(userData.Name)
	go s.mailjetWrapper.SendEmail(ctx, userData.Email, "Konfirmasi Email Berhasil", emailBody)
	s.cache.DeleteEmailVerification(ctx, req.Email)

	return
}