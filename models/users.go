package models

import (
	"database/sql/driver"
	"errors"
	"fmt"

	"gorm.io/gorm"
)


type UserRole string

// Define the allowed roles
const (
	RoleAdmin      UserRole = "admin"
	RoleSales      UserRole = "sales"
	RoleShopOwner  UserRole = "shop_owner"
	RoleShopKeeper UserRole = "shop_keeper"
)

// String method (useful for logging and JSON)
func (r UserRole) String() string {
	return string(r)
}

// Validate checks if the role is valid
func (r UserRole) IsValid() bool {
	switch r {
	case RoleAdmin, RoleSales, RoleShopOwner, RoleShopKeeper:
		return true
	}
	return false
}

// Value implements driver.Valuer interface (for database)
func (r UserRole) Value() (driver.Value, error) {
	if !r.IsValid() {
		return nil, fmt.Errorf("invalid role: %s", r)
	}
	return string(r), nil
}

// Scan implements sql.Scanner interface
func (r *UserRole) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("invalid role type")
	}

	*r = UserRole(str)
	if !r.IsValid() {
		return fmt.Errorf("invalid role value: %s", str)
	}
	return nil
}

// User Model
type User struct {
	gorm.Model
	FirstName        string    `gorm:"not null"`
	LastName         string    `gorm:"not null"`
	Email            string    `gorm:"uniqueIndex;not null"`
	Password         string    `json:"-" gorm:"not null"`
	Role             UserRole  `gorm:"type:user_role;default:'admin';not null"`
	IsActive         bool      `gorm:"default:true"`
	OrganizationId   uint
	ProfileCompleted bool `gorm:"default:false"`
}