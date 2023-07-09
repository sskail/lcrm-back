package models

import (
	"gorm.io/gorm"
	"time"
)

type Project struct {
	gorm.Model
	Title       string    `gorm:"not null" json:"title"`
	Description string    `gorm:"not null" json:"description"`
	Deadline    time.Time `gorm:"" json:"deadline"`
	Creator     User      `gorm:"foreignKey:CreatorId"`
	CreatorId   uint      `gorm:"not null;"`
	BoardId     uint      //`gorm:"not null;"`
	Users       []*User   `gorm:"many2many:project_members;"`
}

func (p *Project) TableName() string {
	return "project"
}

func AddNewMember(p *Project, u *User) error {
	return DB.Model(p).Association("Users").Append(u)
}

func RemoveMember(p *Project, u *User) error {
	return DB.Model(p).Association("Users").Delete(u)
}

func GetAllProject(p *[]Project) error {
	if err := DB.Preload("Users").Find(p).Error; err != nil {
		return err
	}
	return nil
}

func AddNewProject(p *Project) error {
	return DB.Create(p).Error
}

func GetProjectById(p *Project, id string) error {
	return DB.Preload("Users").Preload("Creator").Find(p, id).Error
}

func DeleteProjectById(p *Project, id string) error {
	return DB.Delete(p, id).Error
}

func UpdateProjectById(p *Project) error {
	return DB.Save(p).Error
}
