package seeds

import (
	"gorm.io/gorm"
	"lcrm2/pkg/seed"
)

func All() []seed.Seed {
	return []seed.Seed{
		seed.Seed{
			Name: "CreateAdminRole",
			Run: func(db *gorm.DB) error {
				return CreateRole(db, "Admin")
			},
		},
		seed.Seed{
			Name: "CreateTeacherRole",
			Run: func(db *gorm.DB) error {
				return CreateRole(db, "Teacher")
			},
		},
		seed.Seed{
			Name: "CreateStudentRole",
			Run: func(db *gorm.DB) error {
				return CreateRole(db, "Student")
			},
		},
	}
}
