package models

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html"
	"lcrm2/utils/token"
	"strings"
)

func GetAllUsers(u *[]User) error {
	if err := DB.Find(u).Error; err != nil {
		return err
	}

	return nil
}

type User struct {
	gorm.Model
	Username  string `gorm:"size:255;not null;unique" json:"username"`
	Password  string `gorm:"size:255;not null;" json:"password"`
	FirstName string `gorm:"size:255;" json:"firstName"`
	LastName  string `gorm:"size:255;" json:"lastName"`
	Email     string `gorm:"size:255;" json:"email"`
	GitApi    string `gorm:"size:255;" json:"gitApi"`
	RoleId    uint   `gorm:"not null;" json:"roleId"`
	//Boards    []Board `gorm:"foreignKey:UserID"`
	Tasks []Task `gorm:"foreignKey:AssignedUserID"`
	//Ratings   []Rating `gorm:"foreignKey:UserID"`
}

func (u *User) TableName() string {
	return "user"
}

func GetUserByID(u *User, id string) error {

	if err := DB.Find(u, id).Error; err != nil {
		return errors.New("User not found!")
	}

	u.PrepareGive()
	return nil
}

func UpdateUser(u *User) error {
	return DB.Save(u).Error
}

func (u *User) PrepareGive() {
	u.Password = ""
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (string, error) {
	var err error
	u := User{}
	err = DB.Model(User{}).Where("username = ?", username).Take(&u).Error
	if err != nil {
		return "", err
	}
	err = VerifyPassword(password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	token, err := token.GenerateToken(u.ID, u.RoleId)
	if err != nil {
		return "", err
	}
	return token, nil

}

func (u *User) SaveUser() (*User, error) {
	fmt.Println("SaveUser")

	var err error
	err = DB.Create(u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("BeforeCreate")
	// turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	// remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil
}
