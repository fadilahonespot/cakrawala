package product

import (
	"context"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/fadilahonespot/cakrawala/domain/entity"
	"github.com/fadilahonespot/cakrawala/usecase/product/model"
	"github.com/fadilahonespot/cakrawala/utils/errors"
	"github.com/fadilahonespot/cakrawala/utils/logger"
	"github.com/fadilahonespot/cakrawala/utils/paginate"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

func (s *defaultProduct) CreateProduct(ctx context.Context, req model.Product) (err error) {
	_, err = s.productRepo.FindProductByCode(ctx, req.Code)
	if err == nil {
		logger.Error(ctx, "product code is exists")
		err = errors.New(http.StatusConflict, "product already exists")
		return
	}

	_, err = s.productTypeRepo.GetProductType(ctx, strconv.Itoa(req.ProductTypeID))
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

	tx := s.productRepo.BeginTrans(ctx)
	dataProduct := entity.Product{
		Name:          req.Name,
		Code:          req.Code,
		IsAvailable:   req.IsAvailable,
		Description:   req.Description,
		ProductTypeID: req.ProductTypeID,
		Weight:        req.Weight,
		Price:         req.Price,
		Stock:         req.Stock,
	}
	err = s.productRepo.Create(ctx, tx, &dataProduct)
	if err != nil {
		tx.Rollback()
		logger.Error(ctx, "failed to create product", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	for i := 0; i < len(req.Images); i++ {
		data := entity.ProductImg{
			Image:     req.Images[i],
			ProductID: dataProduct.ID,
		}
		err = s.productImgRepo.UpdateProductImg(ctx, tx, &data)
		if err != nil {
			tx.Rollback()
			logger.Error(ctx, "failed to update product images", err.Error())
			err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
	}

	tx.Commit()
	return
}

func (s *defaultProduct) GetAllProduct(ctx context.Context, param paginate.Pagination) (resp []model.GetAllProductResponse, count int64, err error) {
	dataProduct, count, err := s.productRepo.FindAll(ctx, param)
	if err != nil {
		logger.Error(ctx, "error fetching all products", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	for i := 0; i < len(dataProduct); i++ {
		data := model.GetAllProductResponse{
			ID:              dataProduct[i].ID,
			Name:            dataProduct[i].Name,
			IsAvailable:     dataProduct[i].IsAvailable,
			ProductTypeID:   dataProduct[i].ProductTypeID,
			ProductTypeName: dataProduct[i].ProductType.Name,
			Price:           dataProduct[i].Price,
			Sold:            dataProduct[i].Sold,
		}
		if len(dataProduct[i].ProductImg) != 0 {
			data.Image = dataProduct[i].ProductImg[0].Image
		}

		resp = append(resp, data)
	}
	return
}

func (s *defaultProduct) GetProductById(ctx context.Context, id int) (resp model.Product, err error) {
	dataProduct, err := s.productRepo.FindByID(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Error(ctx, "product not found", err.Error())
			err = errors.New(http.StatusNotFound, err.Error())
			return
		}
		logger.Error(ctx, "failed to get product", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	resp = model.Product{
		ID:              dataProduct.ID,
		Name:            dataProduct.Name,
		Code:            dataProduct.Code,
		IsAvailable:     dataProduct.IsAvailable,
		Description:     dataProduct.Description,
		ProductTypeID:   dataProduct.ProductTypeID,
		ProductTypeName: dataProduct.ProductType.Name,
		Weight:          dataProduct.Weight,
		Stock:           dataProduct.Stock,
		Price:           dataProduct.Price,
		Sold:            dataProduct.Sold,
	}
	for i := 0; i < len(dataProduct.ProductImg); i++ {
		resp.Images = append(resp.Images, dataProduct.ProductImg[i].Image)
	}
	return
}

func (s *defaultProduct) UpdateProduct(ctx context.Context, id int, req model.Product) (err error) {
	productData, err := s.productRepo.FindByID(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Error(ctx, "product not found", err.Error())
			err = errors.New(http.StatusNotFound, err.Error())
			return
		}
		logger.Error(ctx, "failed to get product", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	if productData.Code != req.Code {
		_, err = s.productRepo.FindProductByCode(ctx, req.Code)
		if err == nil {
			logger.Error(ctx, "product code has been used")
			err = errors.New(http.StatusConflict, "product code has been used")
			return
		}
	}

	productData.Code = req.Code
	productData.IsAvailable = req.IsAvailable
	productData.Description = req.Description
	productData.Name = req.Name
	productData.Price = req.Price
	productData.ProductTypeID = req.ProductTypeID
	productData.Stock = req.Stock
	productData.Weight = req.Weight

	var productImage []entity.ProductImg
	for i := 0; i < len(req.Images); i++ {
		productImage = append(productImage, entity.ProductImg{
			Image: req.Images[i],
		})
	}
	productData.ProductImg = productImage

	tx := s.productRepo.BeginTrans(ctx)
	err = s.productImgRepo.DeleteProductImgByProductId(ctx, tx, strconv.Itoa(productData.ID))
	if err != nil {
		tx.Rollback()
		logger.Error(ctx, "error deleting product image", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	err = s.productRepo.Update(ctx, tx, productData)
	if err != nil {
		tx.Rollback()
		logger.Error(ctx, "error updating product", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	tx.Commit()

	return
}

func (s *defaultProduct) DeleteProduct(ctx context.Context, productId string) (err error) {
	productData, err := s.productRepo.FindByID(ctx, cast.ToInt(productId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Error(ctx, "product not found", err.Error())
			err = errors.New(http.StatusNotFound, err.Error())
			return
		}
		logger.Error(ctx, "failed to get product", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	tx := s.productRepo.BeginTrans(ctx)
	err = s.productRepo.Delete(ctx, tx, productData.ID)
	if err != nil {
		tx.Rollback()
		logger.Error(ctx, "failed delete product", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	tx.Commit()

	return
}

func (s *defaultProduct) UploadProductImage(ctx context.Context, images []*multipart.FileHeader) (resp model.UploadImageResponse, err error) {
	pathFolder := "/product/images/"
	for k := 0; k < len(images); k++ {
		link, errRes := s.fileUpload.Uplaod(ctx, images[k], pathFolder)
		if errRes != nil {
			logger.Error(ctx, "error upload image", errRes.Error())
			err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
		resp.Images = append(resp.Images, link)
	}
	return
}