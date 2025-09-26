package models

import "gorm.io/gorm"

const roleTableName = "roles"

type Role struct {
	gorm.Model
	Name        string       `json:"name" gorm:"not null"`
	Description string       `json:"description"`
	IsActive    bool         `json:"is_active" gorm:"default:true"`
	SchoolID    uint         `json:"school_id" gorm:"not null"`
	School      School       `json:"school" gorm:"foreignKey:SchoolID"`
	Permissions []Permission `json:"permissions" gorm:"many2many:role_permissions;"`
}

type RoleDTO struct {
	DTO
	Name        string `json:"name" validate:"required,min=3,max=50"`
	Description string `json:"description" validate:"omitempty,max=200"`
	IsActive    *bool  `json:"is_active"`
	SchoolID    uint   `json:"school_id" validate:"required"`
}

type RolePage struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

func (RoleDTO) TableName() string {
	return roleTableName
}
