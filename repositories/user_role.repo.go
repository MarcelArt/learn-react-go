
package repositories

import (
	"github.com/MarcelArt/learn-react-go/models"
	"gorm.io/gorm"
)

const userRolePageQuery = "-- Write your query here --"

type IUserRoleRepo interface {
	IBaseCrudRepo[models.UserRole, models.UserRoleDTO, models.UserRolePage]
}

type UserRoleRepo struct {
	BaseCrudRepo[models.UserRole, models.UserRoleDTO, models.UserRolePage]
}

func NewUserRoleRepo(db *gorm.DB) *UserRoleRepo {
	return &UserRoleRepo{
		BaseCrudRepo: BaseCrudRepo[models.UserRole, models.UserRoleDTO, models.UserRolePage]{
			db:        db,
			pageQuery: userRolePageQuery,
		},
	}
}
