
package models

import "gorm.io/gorm"

const userRoleTableName = "user_roles"

type UserRole struct {
	gorm.Model
	UserID  uint  `json:"user_id" gorm:"not null"`
	RoleID  uint  `json:"role_id" gorm:"not null"`
	User    User  `json:"user" gorm:"foreignKey:UserID"`
	Role    Role  `json:"role" gorm:"foreignKey:RoleID"`
}

type UserRoleDTO struct {
	DTO
	UserID uint `json:"user_id" validate:"required"`
	RoleID uint `json:"role_id" validate:"required"`
}

type UserRolePage struct {
	UserID uint `json:"user_id"`
	RoleID uint `json:"role_id"`
}

func (UserRoleDTO) TableName() string {
	return userRoleTableName
}

