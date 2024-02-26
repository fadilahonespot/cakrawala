package handler

import (
	"net/http"

	"github.com/fadilahonespot/cakrawala/domain/entity"
	"github.com/fadilahonespot/cakrawala/usecase/product"
	"github.com/fadilahonespot/cakrawala/usecase/product/model"
	"github.com/fadilahonespot/cakrawala/utils"
	"github.com/fadilahonespot/cakrawala/utils/errors"
	"github.com/fadilahonespot/cakrawala/utils/logger"
	"github.com/fadilahonespot/cakrawala/utils/paginate"
	"github.com/fadilahonespot/cakrawala/utils/response"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

type ProductHandler struct {
	productService product.ProductService
}

func NewProductHandler(productService product.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) AddProduct(c echo.Context) (err error) {
	ctx := c.Request().Context()

	clams := utils.GetClamsJwt(c)
	if clams.Role != string(entity.Admin) {
		logger.Error(ctx, "Access for admin only")
		err := errors.New(http.StatusUnauthorized, "Access for admin only")
		return err
	}
	var req model.Product
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

	err = h.productService.CreateProduct(ctx, req)
	if err != nil {
		return err
	}

	resp := response.ResponseSuccess(nil)
	return c.JSON(http.StatusOK, resp)
}

func (h *ProductHandler) GetAllProduct(c echo.Context) (err error) {
	ctx := c.Request().Context()
	params, err := paginate.GetParams(c)
	if err != nil {
		logger.Error(ctx, "error getting params", err.Error())
		err = errors.New(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	data, count, err := h.productService.GetAllProduct(ctx, params)
	if err != nil {
		return err
	}

	resp := response.HandleSuccessWithPagination(float64(count), params, data)
	return c.JSON(http.StatusOK, resp)
}

func (h *ProductHandler) UpdateProduct(c echo.Context) (err error) {
	ctx := c.Request().Context()

	clams := utils.GetClamsJwt(c)
	if clams.Role != string(entity.Admin) {
		logger.Error(ctx, "Access for admin only")
		err := errors.New(http.StatusUnauthorized, "Access for admin only")
		return err
	}

	productId := cast.ToInt(c.Param("productId"))
	var req model.Product
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

	err = h.productService.UpdateProduct(ctx, productId, req)
	if err != nil {
		return err
	}

	resp := response.ResponseSuccess(nil)
	return c.JSON(http.StatusOK, resp)
}

func (h *ProductHandler) DeleteProduct(c echo.Context) (err error) {
	ctx := c.Request().Context()
	productId := c.Param("productId")

	clams := utils.GetClamsJwt(c)
	if clams.Role != string(entity.Admin) {
		logger.Error(ctx, "Access for admin only")
		err := errors.New(http.StatusUnauthorized, "Access for admin only")
		return err
	}

	err = h.productService.DeleteProduct(ctx, productId)
	if err != nil {
		return err
	}

	resp := response.ResponseSuccess(nil)
	return c.JSON(http.StatusOK, resp)
}

func (h *ProductHandler) GetProductById(c echo.Context) (err error) {
	ctx := c.Request().Context()
	productId := c.Param("productId")

	data, err := h.productService.GetProductById(ctx, cast.ToInt(productId))
	if err != nil {
		return err
	}

	resp := response.ResponseSuccess(data)
	return c.JSON(http.StatusOK, resp)
}

func (h *ProductHandler) UploadProductImages(c echo.Context) (err error) {
	ctx := c.Request().Context()
	clams := utils.GetClamsJwt(c)
	if clams.Role != string(entity.Admin) {
		logger.Error(ctx, "Access for admin only")
		err := errors.New(http.StatusUnauthorized, "Access for admin only")
		return err
	}

	file, err := c.MultipartForm()
	if err != nil {
		logger.Error(ctx, "error uploading product images", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	images := file.File["images"]
	if len(images) == 0 {
		logger.Error(ctx, "image is empty")
		err = errors.New(http.StatusBadRequest, "image is empty")
		return
	}
	if len(images) > 5 {
		logger.Error(ctx, "maximum of 5 images permitted")
		err = errors.New(http.StatusBadRequest, "maximum of 5 images permitted")
		return
	}
	for i := 0; i < len(images); i++ {
		err = utils.ValidationImages(images[i].Filename, int(images[i].Size))
		if err != nil {
			logger.Error(ctx, "error validate file name", err.Error())
			err = errors.New(http.StatusBadRequest, err.Error())
			return
		}
	}

	data, err := h.productService.UploadProductImage(ctx, images)
	if err != nil {
		return err
	}

	resp := response.ResponseSuccess(data)
	return c.JSON(http.StatusOK, resp)
}
