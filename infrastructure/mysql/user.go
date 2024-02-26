package mysql

import (
	"context"

	"github.com/fadilahonespot/cakrawala/domain/entity"
	"github.com/fadilahonespot/cakrawala/domain/repository"
	"gorm.io/gorm"
)

type defaultUserRepo struct {
	db *gorm.DB
}

func SetupUserRepository(db *gorm.DB) repository.UserRepository {
	return &defaultUserRepo{
        db: db,
    }
}

func (r *defaultUserRepo) Create(ctx context.Context, user *entity.User) error {
    return r.db.WithContext(ctx).Create(user).Error
}

func (r *defaultUserRepo) FindAll(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
    err := r.db.WithContext(ctx).Find(&users).Error
    return users, err
}

func (r *defaultUserRepo) FindByID(ctx context.Context, id int) (*entity.User, error) {
	var user entity.User
    err := r.db.WithContext(ctx).Where("id =?", id).First(&user).Error
    return &user, err
}

func (r *defaultUserRepo) Update(ctx context.Context, user *entity.User) error {
    return r.db.WithContext(ctx).Save(user).Error
}

func (r *defaultUserRepo) Delete(ctx context.Context, id int) error {
    return r.db.WithContext(ctx).Delete(&entity.User{}, id).Error
}

func (r *defaultUserRepo) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
    err := r.db.WithContext(ctx).Where("email =?", email).First(&user).Error
    return &user, err
}

