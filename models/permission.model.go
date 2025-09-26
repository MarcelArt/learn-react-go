
package models

import "gorm.io/gorm"

const permissionTableName = "permissions"

type Permission struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null;unique"`
	Description string `json:"description"`
	Resource    string `json:"resource" gorm:"not null"`
	Action      string `json:"action" gorm:"not null"`
	IsActive    bool   `json:"is_active" gorm:"default:true"`
}

type PermissionDTO struct {
	DTO
	Name        string `json:"name" validate:"required,min=3,max=50"`
	Description string `json:"description" validate:"omitempty,max=200"`
	Resource    string `json:"resource" validate:"required,min=3,max=50"`
	Action      string `json:"action" validate:"required,min=3,max=50"`
	IsActive    *bool  `json:"is_active"`
}

type PermissionPage struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Resource    string `json:"resource"`
	Action      string `json:"action"`
	IsActive    bool   `json:"is_active"`
}

func (PermissionDTO) TableName() string {
	return permissionTableName
}

