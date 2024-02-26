package product

import (
	"context"
	"mime/multipart"

	"github.com/fadilahonespot/cakrawala/usecase/product/model"
	"github.com/fadilahonespot/cakrawala/utils/paginate"
)

type ProductService interface {
    CreateProductType(ctx context.Context, req model.ProductTypeRequest) (err error)
	UpdateProductType(ctx context.Context, productTypeId string, req model.ProductTypeRequest) (err error)
	GetAllProductTypes(ctx context.Context) (resp []model.GetAllProductTypeResponse, err error)
	CreateProduct(ctx context.Context, req model.Product) (err error)
	GetAllProduct(ctx context.Context, params paginate.Pagination) (resp []model.GetAllProductResponse, count int64, err error)
	UpdateProduct(ctx context.Context, id int, req model.Product) (err error)
	DeleteProduct(ctx context.Context, productId string) (err error)
	GetProductById(ctx context.Context, id int) (resp model.Product, err error)
	UploadProductImage(ctx context.Context, images []*multipart.FileHeader) (resp model.UploadImageResponse, err error)
}