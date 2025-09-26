
package models

import "gorm.io/gorm"

const rolePermissionTableName = "role_permissions"

type RolePermission struct {
	gorm.Model
	RoleID       uint       `json:"role_id" gorm:"not null"`
	PermissionID uint       `json:"permission_id" gorm:"not null"`
	Role         Role       `json:"role" gorm:"foreignKey:RoleID"`
	Permission   Permission `json:"permission" gorm:"foreignKey:PermissionID"`
}

type RolePermissionDTO struct {
	DTO
	RoleID       uint `json:"role_id" validate:"required"`
	PermissionID uint `json:"permission_id" validate:"required"`
}

type RolePermissionPage struct {
	RoleID       uint `json:"role_id"`
	PermissionID uint `json:"permission_id"`
}

func (RolePermissionDTO) TableName() string {
	return rolePermissionTableName
}

