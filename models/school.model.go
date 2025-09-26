
package models

import "gorm.io/gorm"

const schoolTableName = "schools"

type School struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null;unique"`
	Email       string `json:"email" gorm:"not null;unique"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	Website     string `json:"website"`
	Description string `json:"description"`
	LogoURL     string `json:"logo_url"`
	IsActive    bool   `json:"is_active" gorm:"default:true"`
}

type SchoolDTO struct {
	DTO
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Email       string `json:"email" validate:"required,email"`
	Phone       string `json:"phone" validate:"omitempty,e164"`
	Address     string `json:"address" validate:"omitempty,max=200"`
	Website     string `json:"website" validate:"omitempty,url"`
	Description string `json:"description" validate:"omitempty,max=500"`
	LogoURL     string `json:"logo_url" validate:"omitempty,url"`
	IsActive    *bool  `json:"is_active"`
}

type SchoolPage struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	IsActive    bool   `json:"is_active"`
}

func (SchoolDTO) TableName() string {
	return schoolTableName
}

