package mysql

import (
	"context"

	"github.com/fadilahonespot/cakrawala/domain/entity"
	"github.com/fadilahonespot/cakrawala/domain/repository"
	"github.com/fadilahonespot/cakrawala/utils/paginate"
	"gorm.io/gorm"
)

type defaultProductRepo struct {
	db *gorm.DB
}

func SetupProductRepoRepository(db *gorm.DB) repository.ProductRepository {
	return &defaultProductRepo{
        db: db,
    }
}

func (repo *defaultProductRepo) Create(ctx context.Context, tx *gorm.DB, product *entity.Product) (err error) {
	err = tx.WithContext(ctx).Create(product).Error
	return
}

func (repo *defaultProductRepo) FindAll(ctx context.Context, param paginate.Pagination) (products []entity.Product, count int64, err error) {
	query := func(condision *gorm.DB) *gorm.DB {
		if param.Value != "" {
			searchValue := "%" + param.Value + "%"
			switch param.Key {
			case "categoryId":
				condision.Where("product_type_id = ?", param.Value)
			case "description":
				condision.Where("description LIKE ?", searchValue)
			case "name":
				condision.Where("name LIKE ?", searchValue)
			case "weight":
				condision.Where("weight <= ?", param.Value)
			case "stock":
				condision.Where("stock <= ?", param.Value)
			case "price":
				condision.Where("price <= ?", param.Value)
			default:
				condision.Where("name LIKE ?", searchValue)
			}
		}
		
		return condision
	}
	err = repo.db.WithContext(ctx).Model(&entity.Product{}).Scopes(query).Count(&count).Error
	if err != nil {
		return
	}
	err = repo.db.WithContext(ctx).Preload("ProductImg").Preload("ProductType").
		Scopes(paginate.Paginate(param.Page, param.Limit)).Scopes(query).Find(&products).Error
    return
}

func (repo *defaultProductRepo) FindByID(ctx context.Context, id int) (product *entity.Product, err error) {
    err = repo.db.WithContext(ctx).Preload("ProductImg").Preload("ProductType").
		Where("id =?", id).First(&product).Error
    return
}

func (repo *defaultProductRepo) Update(ctx context.Context, tx *gorm.DB, product *entity.Product) (err error) {
    err = tx.WithContext(ctx).Save(product).Error
    return
}

func (repo *defaultProductRepo) Delete(ctx context.Context, tx *gorm.DB, id int) (err error) {
    err = tx.WithContext(ctx).Delete(&entity.Product{}, id).Error
    return
}

func (repo *defaultProductRepo) BeginTrans(ctx context.Context) *gorm.DB {
	return repo.db.Begin()
}
func (repo *defaultProductRepo) FindProductByCode(ctx context.Context, code string) (resp *entity.Product, err error) {
	err = repo.db.WithContext(ctx).Take(&resp, "code = ?", code).Error
	return
}