package handler

import (
	"net/http"
	"strconv"

	"github.com/fadilahonespot/cakrawala/usecase/user"
	"github.com/fadilahonespot/cakrawala/usecase/user/model"
	"github.com/fadilahonespot/cakrawala/utils"
	"github.com/fadilahonespot/cakrawala/utils/constrans"
	"github.com/fadilahonespot/cakrawala/utils/errors"
	"github.com/fadilahonespot/cakrawala/utils/logger"
	"github.com/fadilahonespot/cakrawala/utils/response"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.UserService
}

func NewUserHandler(userService user.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Login(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var req model.LoginRequest
	err = c.Bind(&req)
	if err != nil {
		logger.Error(ctx, "error binding", err.Error())
		err = errors.New(http.StatusBadRequest, err.Error())
		return
	}

	err = c.Validate(req)
	if err != nil {
		logger.Error(ctx, "error validating", err.Error())
		err = errors.New(http.StatusBadRequest, err.Error())
		return
	}

	logger.Info(ctx, "[Request]", req)

	token, err := h.userService.Login(ctx, req)
	if err != nil {
		return err
	}

	resp := response.ResponseSuccess(token)
	return c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) Register(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var req model.RegisterRequest
	err = c.Bind(&req)
	if err != nil {
		logger.Error(ctx, "error binding", err.Error())
		err = errors.New(http.StatusBadRequest, err.Error())
		return
	}

	err = c.Validate(req)
	if err != nil {
		logger.Error(ctx, "error validating", err.Error())
		err = errors.New(http.StatusBadRequest, err.Error())
		return
	}

	logger.Info(ctx, "[Request]", req)

	err = h.userService.Register(ctx, req)
	if err != nil {
		return err
	}

	resp := response.ResponseSuccessWithMessage(http.StatusOK, constrans.SuccessSendEmail)
	return c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) GetProvince(c echo.Context) (err error) {
	ctx := c.Request().Context()

	provinceData, err := h.userService.GetProvince(ctx)
	if err != nil {
		return err
	}

	resp := response.ResponseSuccess(provinceData)
	return c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) GetCity(c echo.Context) (err error) {
	ctx := c.Request().Context()
	provinceId := c.Param("provinceId")

	provinceData, err := h.userService.GetCity(ctx, provinceId)
	if err != nil {
		return err
	}

	resp := response.ResponseSuccess(provinceData)
	return c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) VerificationEmail(c echo.Context) (err error) {
	ctx := c.Request().Context()
	email := c.QueryParam("email")
	id := c.QueryParam("id")

	req := model.VerificationRequest{
		Email: email,
		Id:    id,
	}

	err = h.userService.VerificationEmail(ctx, req)
	if err != nil {
		return err
	}

	resp := response.ResponseSuccess(nil)
	return c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) ResendEmailVerification(c echo.Context) (err error) {
	ctx := c.Request().Context()

	clamsData := utils.GetClamsJwt(c)
	err = h.userService.ResendEmail(ctx, strconv.Itoa(int(clamsData.UserId)))
	if err != nil {
		return err
	}

	resp := response.ResponseSuccessWithMessage(http.StatusOK, constrans.SuccessSendEmail)
	return c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) GetProfile(c echo.Context) (err error) {
	ctx := c.Request().Context()
	clamsData := utils.GetClamsJwt(c)
	req := model.GetProfileRequest{
		UserId:   strconv.Itoa(int(clamsData.UserId)),
		RoleName: clamsData.Role,
	}

	provinceData, err := h.userService.GetProfile(ctx, req)
	if err != nil {
		return err
	}

	resp := response.ResponseSuccess(provinceData)
	return c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) UpdateProfile(c echo.Context) (err error) {
	ctx := c.Request().Context()
	clamsData := utils.GetClamsJwt(c)
	var req model.UpdateProfileRequest
	err = c.Bind(&req)
	if err != nil {
		logger.Error(ctx, "error binding", err.Error())
		err = errors.New(http.StatusBadRequest, err.Error())
		return
	}

	err = c.Validate(req)
	if err != nil {
		logger.Error(ctx, "error validating", err.Error())
		err = errors.New(http.StatusBadRequest, err.Error())
		return
	}

	logger.Info(ctx, "[Request]", req)

	err = h.userService.UpdateProfile(ctx, int(clamsData.UserId), req)
	if err != nil {
		return err
	}

	resp := response.ResponseSuccess(nil)
	return c.JSON(http.StatusOK, resp)
}
