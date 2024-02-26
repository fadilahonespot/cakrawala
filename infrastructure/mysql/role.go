package mysql

// import (
// 	"context"

// 	"github.com/fadilahonespot/first-store/domain/entity"
// 	"github.com/fadilahonespot/first-store/domain/repository"
// 	"gorm.io/gorm"
// )

// type defaultRole struct {
// 	db *gorm.DB
// }

// func SetupRoleRepository(db *gorm.DB) repository.RoleRepository {
// 	return &defaultRole{
//         db: db,
//     }
// }

// func (r *defaultRole) GetRole(ctx context.Context, id string) (*entity.Role, error) {
// 	var resp entity.Role
//     err := r.db.WithContext(ctx).First(&resp, "id =?", id).Error
//     return &resp, err
// }