
package repositories

import (
	"github.com/MarcelArt/learn-react-go/models"
	"gorm.io/gorm"
)

const rolePermissionPageQuery = "-- Write your query here --"

type IRolePermissionRepo interface {
	IBaseCrudRepo[models.RolePermission, models.RolePermissionDTO, models.RolePermissionPage]
}

type RolePermissionRepo struct {
	BaseCrudRepo[models.RolePermission, models.RolePermissionDTO, models.RolePermissionPage]
}

func NewRolePermissionRepo(db *gorm.DB) *RolePermissionRepo {
	return &RolePermissionRepo{
		BaseCrudRepo: BaseCrudRepo[models.RolePermission, models.RolePermissionDTO, models.RolePermissionPage]{
			db:        db,
			pageQuery: rolePermissionPageQuery,
		},
	}
}
