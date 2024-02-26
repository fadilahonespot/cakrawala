package product

import (
	"github.com/fadilahonespot/cakrawala/domain/repository"
	"github.com/fadilahonespot/cakrawala/infrastructure/cached"
	"github.com/fadilahonespot/cakrawala/infrastructure/http/filebox"
)

type defaultProduct struct {
	productRepo     repository.ProductRepository
	productTypeRepo repository.ProductTypeRepository
	productImgRepo  repository.ProductImgRepository
	fileUpload      filebox.DropboxWrapper
	cache           cached.RedisClient
}

func SetupProductService() *defaultProduct {
	return &defaultProduct{}
}

func (s *defaultProduct) SetProductRepo(repo repository.ProductRepository) *defaultProduct {
	s.productRepo = repo
	return s
}

func (s *defaultProduct) SetProductTypeRepo(repo repository.ProductTypeRepository) *defaultProduct {
	s.productTypeRepo = repo
	return s
}

func (s *defaultProduct) SetProductImgRepo(repo repository.ProductImgRepository) *defaultProduct {
	s.productImgRepo = repo
	return s
}

func (s *defaultProduct) SetDropBoxWrapper(wrapper filebox.DropboxWrapper) *defaultProduct {
	s.fileUpload = wrapper
	return s
}

func (s *defaultProduct) SetCache(cache cached.RedisClient) *defaultProduct {
	s.cache = cache
	return s
}

func (s *defaultProduct) Validate() ProductService {
	if s.productRepo == nil {
		panic("product repo is nil")
	}

	if s.productTypeRepo == nil {
		panic("product type repo is nil")
	}

	if s.productImgRepo == nil {
		panic("product image repo is nil")
	}

	if s.fileUpload == nil {
		panic("dropbox wrapper is nil")
	}

	if s.cache == nil {
		panic("cache is nil")
	}

	return s
}
