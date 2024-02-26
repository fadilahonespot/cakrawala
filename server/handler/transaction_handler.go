package handler

import (
	"net/http"
	"os"
	"strconv"

	"github.com/fadilahonespot/cakrawala/usecase/transaction"
	"github.com/fadilahonespot/cakrawala/usecase/transaction/model"
	"github.com/fadilahonespot/cakrawala/utils"
	"github.com/fadilahonespot/cakrawala/utils/errors"
	"github.com/fadilahonespot/cakrawala/utils/logger"
	"github.com/fadilahonespot/cakrawala/utils/paginate"
	"github.com/fadilahonespot/cakrawala/utils/response"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

type TransactionHandler struct {
	transactionService transaction.TransactionService
}

func NewTransactionHandler(transactionService transaction.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		transactionService: transactionService,
	}
}

func (h *TransactionHandler) CheckAvailableBank(c echo.Context) (err error) {
	ctx := c.Request().Context()

	data, err := h.transactionService.CheckAvailableBank(ctx)
	if err != nil {
		return err
	}

	resp := response.ResponseSuccess(data)
	return c.JSON(http.StatusOK, resp)
}

func (h *TransactionHandler) CheckAvailabeCourier(c echo.Context) (err error) {
	ctx := c.Request().Context()

	data, err := h.transactionService.GetCourierInfo(ctx)
	if err != nil {
		return err
	}

	resp := response.ResponseSuccess(data)
	return c.JSON(http.StatusOK, resp)
}

func (h *TransactionHandler) CheckShipping(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var req model.CheckShippingRequest
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

	clams := utils.GetClamsJwt(c)
	ongkirData, err := h.transactionService.CheckShipping(ctx, strconv.Itoa(int(clams.UserId)), req)
	if err != nil {
		return err
	}

	resp := response.ResponseSuccess(ongkirData)
	return c.JSON(http.StatusOK, resp)
}

func (h *TransactionHandler) GetBasket(c echo.Context) (err error) {
	ctx := c.Request().Context()
	clamsData := utils.GetClamsJwt(c)
	data, err := h.transactionService.GetAllProductBasket(ctx, cast.ToString(clamsData.UserId))
	if err != nil {
		return err
	}

	resp := response.ResponseSuccess(data)
	return c.JSON(http.StatusOK, resp)
}

func (h *TransactionHandler) AddBasket(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var req model.AddBasketRequest
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

	clams := utils.GetClamsJwt(c)
	err = h.transactionService.AddProductBasket(ctx, strconv.Itoa(int(clams.UserId)), req)
	if err != nil {
		return err
	}

	resp := response.ResponseSuccess(nil)
	return c.JSON(http.StatusOK, resp)
}

func (h *TransactionHandler) DeleteBasket(c echo.Context) (err error) {
	ctx := c.Request().Context()
	clamsData := utils.GetClamsJwt(c)
	productId := c.Param("productId")
	err = h.transactionService.DeleteProductBasket(ctx, cast.ToString(clamsData.UserId), productId)
	if err != nil {
		return err
	}

	resp := response.ResponseSuccess(nil)
	return c.JSON(http.StatusOK, resp)
}

func (h *TransactionHandler) CheckoutTransaction(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var req model.CheckoutRequest
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

	clams := utils.GetClamsJwt(c)
	dataResp, err := h.transactionService.Checkout(ctx, strconv.Itoa(int(clams.UserId)), req)
	if err != nil {
		return err
	}

	resp := response.ResponseSuccess(dataResp)
	return c.JSON(http.StatusOK, resp)
}

func (h *TransactionHandler) CallbackTransaction(c echo.Context) (err error) {
	ctx := c.Request().Context()
	webHookToken := c.Request().Header.Get("x-callback-token")
	if webHookToken != os.Getenv("XENDIT_WEBHOOK_TOKEN") {
		logger.Error(ctx, "Xendit callback token not valid")
		err = errors.New(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	var req model.CallbackRequest
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

	err = h.transactionService.CallbackTransaction(ctx, "", req)
	if err != nil {
		return err
	}

	resp := response.ResponseSuccess(nil)
	return c.JSON(http.StatusOK, resp)
}

func (h *TransactionHandler) GetHistory(c echo.Context) (err error) {
	ctx := c.Request().Context()
	clamsData := utils.GetClamsJwt(c)
	params, err := paginate.GetParams(c)
	if err != nil {
		logger.Error(ctx, "error getting params", err.Error())
		err = errors.New(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	data, count, err := h.transactionService.GetHistoryTransactio(ctx, cast.ToString(clamsData.UserId), params)
	if err != nil {
		return err
	}

	resp := response.HandleSuccessWithPagination(float64(count), params, data)
	return c.JSON(http.StatusOK, resp)
}

func (h *TransactionHandler) GetDetailTransaction(c echo.Context) (err error) {
	ctx := c.Request().Context()
	trxId := c.Param("transactionId")
	data, err := h.transactionService.GetDetailTransaction(ctx, trxId)
	if err != nil {
		return err
	}

	resp := response.ResponseSuccess(data)
	return c.JSON(http.StatusOK, resp)
}
