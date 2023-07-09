package seeds

import (
	"gorm.io/gorm"
	"lcrm2/models"
)

func CreateRole(db *gorm.DB, name string) error {
	return db.Create(&models.Role{Name: name}).Error
}
