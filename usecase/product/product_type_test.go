package product

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/fadilahonespot/cakrawala/domain/entity"
	"github.com/fadilahonespot/cakrawala/usecase/product/model"
	"github.com/fadilahonespot/cakrawala/utils/logger"
	"github.com/fadilahonespot/cakrawala/utils/mocks"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func Test_defaultProduct_CreateProductType(t *testing.T) {
	logger.NewLogger()
	ctx := context.Background()

	type args struct {
		ctx context.Context
		req model.ProductTypeRequest
	}
	tests := []struct {
		name          string
		args          args
		createTypeErr error
		wantErr       bool
	}{
		{
			name: "failed to create product type",
			args: args{
				ctx: ctx,
				req: model.ProductTypeRequest{
					Name: "makanan",
				},
			},
			createTypeErr: errors.New("error creating product type"),
			wantErr:       true,
		},
		{
			name: "success to create product type",
			args: args{
				ctx: ctx,
				req: model.ProductTypeRequest{
					Name: "makanan",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productTypeRepo := new(mocks.ProductTypeRepository)
			productTypeRepo.On("CreateProductType", mock.Anything, mock.Anything).Return(tt.createTypeErr).Once()

			s := SetupProductService().SetProductTypeRepo(productTypeRepo)
			if err := s.CreateProductType(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("defaultProduct.CreateProductType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_defaultProduct_UpdateProductType(t *testing.T) {
	logger.NewLogger()
	ctx := context.Background()

	type args struct {
		ctx           context.Context
		productTypeId string
		req           model.ProductTypeRequest
	}
	tests := []struct {
		name               string
		args               args
		getProductTypeResp *entity.ProductType
		getProductTypeErr  error
		updateTypeErr      error
		wantErr            bool
	}{
		{
			name: "product type not found",
			args: args{
				ctx:           ctx,
				productTypeId: "1",
				req: model.ProductTypeRequest{
					Name: "makanan",
				},
			},
			getProductTypeErr: gorm.ErrRecordNotFound,
			wantErr:           true,
		},
		{
			name: "failed to get product type",
			args: args{
				ctx:           ctx,
				productTypeId: "1",
				req: model.ProductTypeRequest{
					Name: "makanan",
				},
			},
			getProductTypeErr: errors.New("error getting product type"),
			wantErr:           true,
		},
		{
			name: "failed to update product type",
			args: args{
				ctx:           ctx,
				productTypeId: "1",
				req: model.ProductTypeRequest{
					Name: "makanan",
				},
			},
			getProductTypeResp: &entity.ProductType{
				ID: 1,
			},
			updateTypeErr: errors.New("error"),
			wantErr:       true,
		},
		{
			name: "success to update product type",
			args: args{
				ctx:           ctx,
				productTypeId: "1",
				req: model.ProductTypeRequest{
					Name: "makanan",
				},
			},
			getProductTypeResp: &entity.ProductType{
				ID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productTypeRepo := new(mocks.ProductTypeRepository)
			productTypeRepo.On("GetProductType", mock.Anything, mock.Anything).Return(tt.getProductTypeResp, tt.getProductTypeErr).Once()
			productTypeRepo.On("UpdateProductType", mock.Anything, mock.Anything).Return(tt.updateTypeErr).Once()

			s := SetupProductService().SetProductTypeRepo(productTypeRepo)
			if err := s.UpdateProductType(tt.args.ctx, tt.args.productTypeId, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("defaultProduct.UpdateProductType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_defaultProduct_GetAllProductTypes(t *testing.T) {
	logger.NewLogger()
	ctx := context.Background()
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name       string
		args       args
		getAllResp []entity.ProductType
		getAllErr  error
		wantResp   []model.GetAllProductTypeResponse
		wantErr    bool
	}{
		{
			name: "failed to get all product types",
			args: args{
				ctx: ctx,
			},
			getAllErr: errors.New("error getting all product types"),
			wantErr: true,
		},
		{
			name: "success to get all product types",
			args: args{
				ctx: ctx,
			},
			getAllResp: []entity.ProductType{
				{
					ID: 1,
					Name: "makanan",
				},
			},
			wantResp: []model.GetAllProductTypeResponse{
				{
					Id: 1,
					Name: "makanan",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productTypeRepo := new(mocks.ProductTypeRepository)
			productTypeRepo.On("GetAllProductTypes", mock.Anything).Return(tt.getAllResp, tt.getAllErr).Once()

			s := SetupProductService().SetProductTypeRepo(productTypeRepo)
			gotResp, err := s.GetAllProductTypes(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("defaultProduct.GetAllProductTypes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("defaultProduct.GetAllProductTypes() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
