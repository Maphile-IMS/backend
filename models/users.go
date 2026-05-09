package models

import "gorm.io/gorm"



type User struct {
	gorm.Model
	FirstName     string `gorm:"not null"`
	LastName      string `gorm:"not null"`
	Email    string `gorm:"unique"`
	Password string `json:"-"` 
	Role    string `gorm:"type:enum('admin','sales','shop_owner','shop_keeper');default:'admin';not null"`
	IsActive bool   `gorm:"default:true"`
	OrganizationId uint
	ProfileCompleted bool `gorm:"default:false"`
}
