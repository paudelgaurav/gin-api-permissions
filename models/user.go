package models

import (
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string         `gorm:"size:255;not null;unique" json:"username" form:"username" binding:"required"`
	Password    string         `gorm:"not null" form:"password" json:"-"`
	Permissions pq.StringArray `gorm:"type:varchar(200)[]" json:"permissions" form:"permissions"`
}

func (u *User) BeforeSave() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
