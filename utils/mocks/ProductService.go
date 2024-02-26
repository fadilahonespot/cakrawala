// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/fadilahonespot/cakrawala/usecase/product/model"
	mock "github.com/stretchr/testify/mock"

	multipart "mime/multipart"

	paginate "github.com/fadilahonespot/cakrawala/utils/paginate"
)

// ProductService is an autogenerated mock type for the ProductService type
type ProductService struct {
	mock.Mock
}

// CreateProduct provides a mock function with given fields: ctx, req
func (_m *ProductService) CreateProduct(ctx context.Context, req model.Product) error {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for CreateProduct")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Product) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateProductType provides a mock function with given fields: ctx, req
func (_m *ProductService) CreateProductType(ctx context.Context, req model.ProductTypeRequest) error {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for CreateProductType")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.ProductTypeRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteProduct provides a mock function with given fields: ctx, productId
func (_m *ProductService) DeleteProduct(ctx context.Context, productId string) error {
	ret := _m.Called(ctx, productId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteProduct")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, productId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GenerateDescriptionProduct provides a mock function with given fields: ctx, productId
func (_m *ProductService) GenerateDescriptionProduct(ctx context.Context, productId string) (model.GenerateTextResponse, error) {
	ret := _m.Called(ctx, productId)

	if len(ret) == 0 {
		panic("no return value specified for GenerateDescriptionProduct")
	}

	var r0 model.GenerateTextResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (model.GenerateTextResponse, error)); ok {
		return rf(ctx, productId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) model.GenerateTextResponse); ok {
		r0 = rf(ctx, productId)
	} else {
		r0 = ret.Get(0).(model.GenerateTextResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, productId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllProduct provides a mock function with given fields: ctx, params
func (_m *ProductService) GetAllProduct(ctx context.Context, params paginate.Pagination) ([]model.GetAllProductResponse, int64, error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for GetAllProduct")
	}

	var r0 []model.GetAllProductResponse
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, paginate.Pagination) ([]model.GetAllProductResponse, int64, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, paginate.Pagination) []model.GetAllProductResponse); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.GetAllProductResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, paginate.Pagination) int64); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, paginate.Pagination) error); ok {
		r2 = rf(ctx, params)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetAllProductTypes provides a mock function with given fields: ctx
func (_m *ProductService) GetAllProductTypes(ctx context.Context) ([]model.GetAllProductTypeResponse, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllProductTypes")
	}

	var r0 []model.GetAllProductTypeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]model.GetAllProductTypeResponse, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []model.GetAllProductTypeResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.GetAllProductTypeResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGenerateProductDesc provides a mock function with given fields: ctx, productId
func (_m *ProductService) GetGenerateProductDesc(ctx context.Context, productId string) (model.GenerateTextResponse, error) {
	ret := _m.Called(ctx, productId)

	if len(ret) == 0 {
		panic("no return value specified for GetGenerateProductDesc")
	}

	var r0 model.GenerateTextResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (model.GenerateTextResponse, error)); ok {
		return rf(ctx, productId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) model.GenerateTextResponse); ok {
		r0 = rf(ctx, productId)
	} else {
		r0 = ret.Get(0).(model.GenerateTextResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, productId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProductById provides a mock function with given fields: ctx, id
func (_m *ProductService) GetProductById(ctx context.Context, id int) (model.Product, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetProductById")
	}

	var r0 model.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (model.Product, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) model.Product); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.Product)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProduct provides a mock function with given fields: ctx, id, req
func (_m *ProductService) UpdateProduct(ctx context.Context, id int, req model.Product) error {
	ret := _m.Called(ctx, id, req)

	if len(ret) == 0 {
		panic("no return value specified for UpdateProduct")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, model.Product) error); ok {
		r0 = rf(ctx, id, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateProductType provides a mock function with given fields: ctx, productTypeId, req
func (_m *ProductService) UpdateProductType(ctx context.Context, productTypeId string, req model.ProductTypeRequest) error {
	ret := _m.Called(ctx, productTypeId, req)

	if len(ret) == 0 {
		panic("no return value specified for UpdateProductType")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, model.ProductTypeRequest) error); ok {
		r0 = rf(ctx, productTypeId, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UploadProductImage provides a mock function with given fields: ctx, images
func (_m *ProductService) UploadProductImage(ctx context.Context, images []*multipart.FileHeader) (model.UploadImageResponse, error) {
	ret := _m.Called(ctx, images)

	if len(ret) == 0 {
		panic("no return value specified for UploadProductImage")
	}

	var r0 model.UploadImageResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []*multipart.FileHeader) (model.UploadImageResponse, error)); ok {
		return rf(ctx, images)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []*multipart.FileHeader) model.UploadImageResponse); ok {
		r0 = rf(ctx, images)
	} else {
		r0 = ret.Get(0).(model.UploadImageResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, []*multipart.FileHeader) error); ok {
		r1 = rf(ctx, images)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewProductService creates a new instance of ProductService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProductService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProductService {
	mock := &ProductService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
