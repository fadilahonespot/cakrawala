package product

import (
	"context"
	"errors"
	"mime/multipart"
	"reflect"
	"testing"

	"github.com/fadilahonespot/cakrawala/domain/entity"
	"github.com/fadilahonespot/cakrawala/usecase/product/model"
	"github.com/fadilahonespot/cakrawala/utils"
	"github.com/fadilahonespot/cakrawala/utils/logger"
	"github.com/fadilahonespot/cakrawala/utils/mocks"
	"github.com/fadilahonespot/cakrawala/utils/paginate"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func Test_defaultProduct_CreateProduct(t *testing.T) {
	db := utils.MockGorm()
	logger.NewLogger()
	ctx := context.Background()

	type args struct {
		ctx context.Context
		req model.Product
	}
	tests := []struct {
		name               string
		args               args
		findByCodeResp     *entity.Product
		findByCodeErr      error
		getProductTypeReps *entity.ProductType
		getProductTypeErr  error
		createProductErr   error
		updateImgErr       error
		wantErr            bool
	}{
		{
			name: "product code is exists",
			args: args{
				ctx: ctx,
				req: model.Product{
					Name: "ciptanden",
					Code: "rty-34223",
				},
			},
			wantErr: true,
		},
		{
			name: "product type not found",
			args: args{
				ctx: ctx,
				req: model.Product{
					Name: "ciptanden",
					Code: "rty-34223",
				},
			},
			findByCodeErr:     errors.New("error"),
			getProductTypeErr: gorm.ErrRecordNotFound,
			wantErr:           true,
		},
		{
			name: "failed to get product type",
			args: args{
				ctx: ctx,
				req: model.Product{
					Name: "ciptanden",
					Code: "rty-34223",
				},
			},
			findByCodeErr:     errors.New("error"),
			getProductTypeErr: errors.New("failed to get product type"),
			wantErr:           true,
		},
		{
			name: "failed create product",
			args: args{
				ctx: ctx,
				req: model.Product{
					Name: "ciptanden",
					Code: "rty-34223",
				},
			},
			findByCodeErr:    errors.New("error"),
			createProductErr: errors.New("error"),
			wantErr:          true,
		},
		{
			name: "failed update product images",
			args: args{
				ctx: ctx,
				req: model.Product{
					Name:   "ciptanden",
					Code:   "rty-34223",
					Images: []string{"http://google.com/gambar.jpeg"},
				},
			},
			findByCodeErr: errors.New("error"),
			updateImgErr:  errors.New("error"),
			wantErr:       true,
		},
		{
			name: "success create product",
			args: args{
				ctx: ctx,
				req: model.Product{
					Name:   "ciptanden",
					Code:   "rty-34223",
					Images: []string{"http://google.com/gambar.jpeg"},
				},
			},
			findByCodeErr: errors.New("error"),
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productRepo := new(mocks.ProductRepository)
			productTypeRepo := new(mocks.ProductTypeRepository)
			productImgRepo := new(mocks.ProductImgRepository)
			fileUploadWrapper := new(mocks.DropboxWrapper)
			cache := new(mocks.RedisClient)

			productRepo.On("FindProductByCode", mock.Anything, mock.Anything).Return(tt.findByCodeResp, tt.findByCodeErr).Once()
			productTypeRepo.On("GetProductType", mock.Anything, mock.Anything).Return(tt.getProductTypeReps, tt.getProductTypeErr).Once()
			productRepo.On("BeginTrans", mock.Anything).Return(db).Once()
			productRepo.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(tt.createProductErr).Once()
			productImgRepo.On("UpdateProductImg", mock.Anything, mock.Anything, mock.Anything).Return(tt.updateImgErr, tt.updateImgErr).Once()

			svc := SetupProductService().
				SetProductRepo(productRepo).
				SetProductTypeRepo(productTypeRepo).
				SetCache(cache).
				SetProductImgRepo(productImgRepo).
				SetDropBoxWrapper(fileUploadWrapper).
				Validate()

			if err := svc.CreateProduct(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("defaultProduct.CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_defaultProduct_GetAllProduct(t *testing.T) {
	logger.NewLogger()
	ctx := context.Background()
	respProduct := []entity.Product{
		{
			ID:   1,
			Name: "sarimi rasa ayam bawang",
		},
	}

	resp := []model.GetAllProductResponse{
		{
			ID:   1,
			Name: "sarimi rasa ayam bawang",
		},
	}

	type args struct {
		ctx   context.Context
		param paginate.Pagination
	}
	tests := []struct {
		name         string
		args         args
		wantResp     []model.GetAllProductResponse
		findAllResp  []entity.Product
		findAllCount int64
		findAllErr   error
		wantCount    int64
		wantErr      bool
	}{
		{
			name: "error fetching all products",
			args: args{
				ctx: ctx,
				param: paginate.Pagination{
					Page:  1,
					Limit: 10,
				},
			},
			findAllErr: errors.New("error fetching all products"),
			wantErr:    true,
		},
		{
			name: "find all products successfully",
			args: args{
				ctx: ctx,
				param: paginate.Pagination{
					Page:  1,
					Limit: 10,
				},
			},
			findAllResp:  respProduct,
			findAllCount: 1,
			wantCount:    1,
			wantResp:     resp,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productRepo := new(mocks.ProductRepository)
			productRepo.On("FindAll", mock.Anything, mock.Anything).Return(tt.findAllResp, tt.findAllCount, tt.findAllErr).Once()

			s := SetupProductService().SetProductRepo(productRepo)
			gotResp, gotCount, err := s.GetAllProduct(tt.args.ctx, tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("defaultProduct.GetAllProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("defaultProduct.GetAllProduct() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
			if gotCount != tt.wantCount {
				t.Errorf("defaultProduct.GetAllProduct() gotCount = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func Test_defaultProduct_GetProductById(t *testing.T) {
	logger.NewLogger()
	ctx := context.Background()
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name     string
		args     args
		findResp *entity.Product
		findErr  error
		wantResp model.Product
		wantErr  bool
	}{
		{
			name: "product not found",
			args: args{
				ctx: ctx,
				id:  1,
			},
			findErr: gorm.ErrRecordNotFound,
			wantErr: true,
		},
		{
			name: "failed to get product",
			args: args{
				ctx: ctx,
				id:  1,
			},
			findErr: errors.New("errors"),
			wantErr: true,
		},
		{
			name: "failed to get product",
			args: args{
				ctx: ctx,
				id:  1,
			},
			findErr: errors.New("errors"),
			wantErr: true,
		},
		{
			name: "success find product",
			args: args{
				ctx: ctx,
				id:  1,
			},
			findResp: &entity.Product{
				ID: 1,
			},
			wantResp: model.Product{
				ID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productRepo := new(mocks.ProductRepository)
			productRepo.On("FindByID", mock.Anything, mock.Anything).Return(tt.findResp, tt.findErr).Once()

			s := SetupProductService().SetProductRepo(productRepo)
			gotResp, err := s.GetProductById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("defaultProduct.GetProductById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("defaultProduct.GetProductById() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func Test_defaultProduct_UpdateProduct(t *testing.T) {
	logger.NewLogger()
	ctx := context.Background()
	db := utils.MockGorm()
	type args struct {
		ctx context.Context
		id  int
		req model.Product
	}
	tests := []struct {
		name                string
		args                args
		findResp            *entity.Product
		findErr             error
		findByCodeResp      *entity.Product
		findByCodeErr       error
		deleteProductImgErr error
		updateErr           error
		wantErr             bool
	}{
		{
			name: "product not found",
			args: args{
				ctx: ctx,
				id:  1,
				req: model.Product{
					ID: 1,
				},
			},
			findErr: gorm.ErrRecordNotFound,
			wantErr: true,
		},
		{
			name: "failed to get product",
			args: args{
				ctx: ctx,
				id:  1,
				req: model.Product{
					ID: 1,
				},
			},
			findErr: errors.New("error"),
			wantErr: true,
		},
		{
			name: "product code has been used",
			args: args{
				ctx: ctx,
				id:  1,
				req: model.Product{
					ID:   1,
					Code: "RT45",
				},
			},
			findResp: &entity.Product{
				Code: "TY67",
			},
			findByCodeResp: &entity.Product{
				Code: "RT45",
			},
			wantErr: true,
		},
		{
			name: "error deleted product image",
			args: args{
				ctx: ctx,
				id:  1,
				req: model.Product{
					ID:   1,
					Code: "RT45",
				},
			},
			findResp: &entity.Product{
				Code: "RT45",
			},
			deleteProductImgErr: errors.New("error deleting product image"),
			wantErr:             true,
		},
		{
			name: "error update product",
			args: args{
				ctx: ctx,
				id:  1,
				req: model.Product{
					ID:   1,
					Code: "RT45",
				},
			},
			findResp: &entity.Product{
				Code: "RT45",
			},
			updateErr: errors.New("error update product"),
			wantErr:   true,
		},
		{
			name: "success update product",
			args: args{
				ctx: ctx,
				id:  1,
				req: model.Product{
					ID:   1,
					Code: "RT45",
				},
			},
			findResp: &entity.Product{
				Code: "RT45",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productRepo := new(mocks.ProductRepository)
			productImgRepo := new(mocks.ProductImgRepository)

			productRepo.On("FindByID", mock.Anything, mock.Anything).Return(tt.findResp, tt.findErr).Once()
			productRepo.On("FindProductByCode", mock.Anything, mock.Anything).Return(tt.findByCodeResp, tt.findByCodeErr).Once()
			productRepo.On("BeginTrans", mock.Anything).Return(db).Once()
			productImgRepo.On("DeleteProductImgByProductId", mock.Anything, mock.Anything, mock.Anything).Return(tt.deleteProductImgErr).Once()
			productRepo.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(tt.updateErr).Once()

			s := SetupProductService().SetProductRepo(productRepo).SetProductImgRepo(productImgRepo)
			if err := s.UpdateProduct(tt.args.ctx, tt.args.id, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("defaultProduct.UpdateProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_defaultProduct_DeleteProduct(t *testing.T) {
	logger.NewLogger()
	ctx := context.Background()
	db := utils.MockGorm()
	type args struct {
		ctx       context.Context
		productId string
	}
	tests := []struct {
		name      string
		args      args
		findResp  *entity.Product
		findErr   error
		deleteErr error
		wantErr   bool
	}{
		{
			name: "product not found",
			args: args{
				ctx:       ctx,
				productId: "1",
			},
			findErr: gorm.ErrRecordNotFound,
			wantErr: true,
		},
		{
			name: "failed to get product",
			args: args{
				ctx:       ctx,
				productId: "1",
			},
			findErr: errors.New("error getting product"),
			wantErr: true,
		},
		{
			name: "failed delete product",
			args: args{
				ctx:       ctx,
				productId: "1",
			},
			findResp: &entity.Product{
				ID: 1,
			},
			deleteErr: errors.New("error deleting product"),
			wantErr:   true,
		},
		{
			name: "success delete product",
			args: args{
				ctx:       ctx,
				productId: "1",
			},
			findResp: &entity.Product{
				ID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productRepo := new(mocks.ProductRepository)
			productRepo.On("FindByID", mock.Anything, mock.Anything).Return(tt.findResp, tt.findErr).Once()
			productRepo.On("BeginTrans", mock.Anything).Return(db).Once()
			productRepo.On("Delete", mock.Anything, mock.Anything, mock.Anything).Return(tt.deleteErr).Once()

			s := SetupProductService().SetProductRepo(productRepo)
			if err := s.DeleteProduct(tt.args.ctx, tt.args.productId); (err != nil) != tt.wantErr {
				t.Errorf("defaultProduct.DeleteProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_defaultProduct_UploadProductImage(t *testing.T) {
	logger.NewLogger()
	ctx := context.Background()
	type args struct {
		ctx    context.Context
		images []*multipart.FileHeader
	}
	tests := []struct {
		name       string
		args       args
		uploadResp string
		uploadErr  error
		wantResp   model.UploadImageResponse
		wantErr    bool
	}{
		{
			name: "error upload image",
			args: args{
				ctx: ctx,
				images: []*multipart.FileHeader{
					{
						Filename: "images-1",
					},
				},
			},
			uploadErr: errors.New("error upload image"),
			wantErr:   true,
		},
		{
			name: "success upload image",
			args: args{
				ctx: ctx,
				images: []*multipart.FileHeader{
					{
						Filename: "images-1",
					},
				},
			},
			uploadResp: "http://google.com/image.jpg",
			wantResp: model.UploadImageResponse{
				Images: []string{"http://google.com/image.jpg"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dropboxWrapper := new(mocks.DropboxWrapper)
			dropboxWrapper.On("Uplaod", mock.Anything, mock.Anything, mock.Anything).Return(tt.uploadResp, tt.uploadErr).Once()

			s := SetupProductService().SetDropBoxWrapper(dropboxWrapper)
			gotResp, err := s.UploadProductImage(tt.args.ctx, tt.args.images)
			if (err != nil) != tt.wantErr {
				t.Errorf("defaultProduct.UploadProductImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("defaultProduct.UploadProductImage() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

