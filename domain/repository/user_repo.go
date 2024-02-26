package repository

import (
	"context"

	"github.com/fadilahonespot/cakrawala/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	FindAll(ctx context.Context) ([]entity.User, error)
	FindByID(ctx context.Context, id int) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id int) error
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
}