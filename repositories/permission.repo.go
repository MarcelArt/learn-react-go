
package repositories

import (
	"github.com/MarcelArt/learn-react-go/models"
	"gorm.io/gorm"
)

const permissionPageQuery = "-- Write your query here --"

type IPermissionRepo interface {
	IBaseCrudRepo[models.Permission, models.PermissionDTO, models.PermissionPage]
}

type PermissionRepo struct {
	BaseCrudRepo[models.Permission, models.PermissionDTO, models.PermissionPage]
}

func NewPermissionRepo(db *gorm.DB) *PermissionRepo {
	return &PermissionRepo{
		BaseCrudRepo: BaseCrudRepo[models.Permission, models.PermissionDTO, models.PermissionPage]{
			db:        db,
			pageQuery: permissionPageQuery,
		},
	}
}
