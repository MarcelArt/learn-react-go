
package repositories

import (
	"github.com/MarcelArt/learn-react-go/models"
	"gorm.io/gorm"
)

const rolePageQuery = "-- Write your query here --"

type IRoleRepo interface {
	IBaseCrudRepo[models.Role, models.RoleDTO, models.RolePage]
}

type RoleRepo struct {
	BaseCrudRepo[models.Role, models.RoleDTO, models.RolePage]
}

func NewRoleRepo(db *gorm.DB) *RoleRepo {
	return &RoleRepo{
		BaseCrudRepo: BaseCrudRepo[models.Role, models.RoleDTO, models.RolePage]{
			db:        db,
			pageQuery: rolePageQuery,
		},
	}
}
