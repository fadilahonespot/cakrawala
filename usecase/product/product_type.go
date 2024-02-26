package product

import (
	"context"
	"net/http"

	"github.com/fadilahonespot/cakrawala/domain/entity"
	"github.com/fadilahonespot/cakrawala/usecase/product/model"
	"github.com/fadilahonespot/cakrawala/utils/errors"
	"github.com/fadilahonespot/cakrawala/utils/logger"
	"gorm.io/gorm"
)

func (s *defaultProduct) CreateProductType(ctx context.Context, req model.ProductTypeRequest) (err error) {
	typeProduct := entity.ProductType{
		Name: req.Name,
	}
	err = s.productTypeRepo.CreateProductType(ctx, &typeProduct)
	if err != nil {
		logger.Error(ctx, "failed to create product type", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	return
}

func (s *defaultProduct) UpdateProductType(ctx context.Context, productTypeId string, req model.ProductTypeRequest) (err error) {
	dataProductType, err := s.productTypeRepo.GetProductType(ctx, productTypeId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Error(ctx, "product type not found", err.Error())
			err = errors.New(http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}
		logger.Error(ctx, "failed to get product type", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	dataProductType.Name = req.Name
	err = s.productTypeRepo.UpdateProductType(ctx, dataProductType)
	if err != nil {
		logger.Error(ctx, "failed to update product type", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	return
}

func (s *defaultProduct) GetAllProductTypes(ctx context.Context) (resp []model.GetAllProductTypeResponse, err error) {
	datas, err := s.productTypeRepo.GetAllProductTypes(ctx)
	if err != nil {
		logger.Error(ctx, "failed to get all product types", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	for i := 0; i < len(datas); i++ {
		resp = append(resp, model.GetAllProductTypeResponse{
            Id:   datas[i].ID,
            Name: datas[i].Name,
        })
	}
	
	return
}
