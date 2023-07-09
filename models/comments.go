package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Body      string `gorm:"size:1024;not null" json:"body"`
	ProjectID uint   `gorm:"not null;" json:"projectID"`
	Creator   User   `gorm:"foreignkey:CreatorId"`
	CreatorId uint   `gorm:"not null;"`
	Replies   []Reply
}

func (p *Comment) TableName() string {
	return "comment"
}

// CreateComment Создание нового комментария
func CreateComment(comment *Comment) error {
	return DB.Create(comment).Error
}

// GetCommentsByProjectID Получение всех комментариев для определенного курса
func GetCommentsByProjectID(projectID string) ([]Comment, error) {
	var comments []Comment
	err := DB.Where("project_id = ?", projectID).Preload("Creator").Preload("Replies.Creator").Find(&comments).Error
	return comments, err
}

type Reply struct {
	gorm.Model
	Body      string `gorm:"size:1024;not null" json:"body"`
	Creator   User   `gorm:"foreignkey:CreatorId"`
	CreatorId uint   `gorm:"not null;"`
	CommentID uint   `gorm:"not null;" json:"commentID"`
}

func (p *Reply) TableName() string {
	return "reply"
}

// CreateReply Создание нового ответа на комментарий
func CreateReply(reply *Reply) error {
	return DB.Create(reply).Error
}

// GetRepliesByCommentID Получение ответов для определенного комментария
func GetRepliesByCommentID(commentID string) ([]Reply, error) {
	var replies []Reply
	err := DB.Where("comment_id = ?", commentID).Find(&replies).Error
	return replies, err
}
