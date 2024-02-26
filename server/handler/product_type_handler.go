package handler

import (
	"net/http"

	"github.com/fadilahonespot/cakrawala/domain/entity"
	"github.com/fadilahonespot/cakrawala/usecase/product"
	"github.com/fadilahonespot/cakrawala/usecase/product/model"
	"github.com/fadilahonespot/cakrawala/utils"
	"github.com/fadilahonespot/cakrawala/utils/errors"
	"github.com/fadilahonespot/cakrawala/utils/logger"
	"github.com/fadilahonespot/cakrawala/utils/response"
	"github.com/labstack/echo/v4"
)

type ProductTypeHandler struct {
	productService product.ProductService
}

func NewProductTypeHandler(productService product.ProductService) *ProductTypeHandler {
	return &ProductTypeHandler{
		productService: productService,
	}
}

func (h *ProductTypeHandler) CreateProductType(c echo.Context) (err error) {
	ctx := c.Request().Context()
	clams := utils.GetClamsJwt(c)
	if clams.Role != string(entity.Admin) {
		logger.Error(ctx, "Access for admin only")
		err := errors.New(http.StatusUnauthorized, "Access for admin only")
		return err
	}
	
	var req model.ProductTypeRequest
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

	err = h.productService.CreateProductType(ctx, req)
	if err != nil {
		return err
	}

	resp := response.ResponseSuccess(nil)
	return c.JSON(http.StatusOK, resp)
}

func (h *ProductTypeHandler) UpdateProductType(c echo.Context) (err error) {
	ctx := c.Request().Context()
	clams := utils.GetClamsJwt(c)
	if clams.Role != string(entity.Admin) {
		logger.Error(ctx, "Access for admin only")
		err := errors.New(http.StatusUnauthorized, "Access for admin only")
		return err
	}
	
	productTypeId := c.Param("productTypeId")
	var req model.ProductTypeRequest
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

	err = h.productService.UpdateProductType(ctx, productTypeId, req)
	if err != nil {
		return err
	}

	resp := response.ResponseSuccess(nil)
	return c.JSON(http.StatusOK, resp)
}

func (h *ProductTypeHandler) GetAllProductType(c echo.Context) (err error) {
	ctx := c.Request().Context()

	data, err := h.productService.GetAllProductTypes(ctx)
	if err != nil {
		return err
	}

	resp := response.ResponseSuccess(data)
	return c.JSON(http.StatusOK, resp)
}

