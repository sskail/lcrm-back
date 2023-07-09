package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name string `gorm:"size:255;not null;unique" json:"name"`
	User []User
}

func (r *Role) TableName() string {
	return "role"
}

func GetRoleById(r *Role, id string) error {
	return DB.Find(r, id).Error
}
