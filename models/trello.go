package models

import "gorm.io/gorm"

type Board struct {
	gorm.Model
	Name        string `gorm:"size:255;not null;unique" json:"name"`
	Description string `gorm:"size:1023;not null" json:"description"`
	//UserID      uint   `gorm:"not null" json:"userID"`
	ProjectID uint   `gorm:"not null" json:"projectID"`
	Tasks     []Task `gorm:"foreignKey:BoardID"`
}

func (b *Board) TableName() string {
	return "board"
}

func CreateBoard(b *Board) error {
	return DB.Create(b).Error
}

func GetBoardById(b *Board, id string) error {
	return DB.Preload("Tasks").Preload("Tasks.User").Find(b, id).Error
	// return storage.DB.Find(a,id).Error
}

type Task struct {
	gorm.Model
	Title          string   `gorm:"size:255;not null;" json:"title"`
	Name           string   `gorm:"size:255;not null;" json:"name"`
	Status         string   `gorm:"size:255;not null;" json:"status"`
	BoardID        uint     `gorm:"not null;" json:"boardID"`
	AssignedUserID uint     `json:"assignedUserID"`
	Ratings        []Rating `gorm:"foreignKey:TaskID"`
	User           User     `gorm:"foreignKey:AssignedUserID"`
}

func CreateTask(t *Task) error {
	return DB.Create(t).Error
}

func DeleteTaskById(t *Task, id string) error {
	return DB.Delete(t, id).Error
}

func GetTaskById(t *Task, id string) error {
	return DB.Preload("Ratings").Find(t, id).Error
	// return storage.DB.Find(a,id).Error
}

func UpdateTaskById(t *Task) error {
	return DB.Save(t).Error
}
func (t *Task) TableName() string {
	return "task"
}

type Rating struct {
	gorm.Model
	Score     int
	CreatorID uint
	UserID    uint
	TaskID    uint
}

func (r *Rating) TableName() string {
	return "rating"
}

func CreateRating(rating *Rating) error {
	return DB.Create(rating).Error
}

func GetRatings(taskID uint) ([]Rating, error) {
	var ratings []Rating
	err := DB.Where("task_id = ?", taskID).Find(&ratings).Error
	return ratings, err
}
